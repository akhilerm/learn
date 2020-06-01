#define	VDEV_UBERBLOCK_RING	(128 << 10)

#define	VDEV_PAD_SIZE		(8 << 10)

typedef struct zio_cksum {
	uint64_t	zc_word[4];
} zio_cksum_t;

typedef struct zio_eck {
	uint64_t	zec_magic;	/* for validation, endianness	*/
	zio_cksum_t	zec_cksum;	/* 256-bit checksum		*/
} zio_eck_t;

typedef struct vdev_phys {
	char		vp_nvlist[VDEV_PHYS_SIZE - sizeof (zio_eck_t)];
	zio_eck_t	vp_zbt;
} vdev_phys_t;


typedef struct vdev_label {
	char		vl_pad1[VDEV_PAD_SIZE];			/*  8K */
	char		vl_pad2[VDEV_PAD_SIZE];			/*  8K */
	vdev_phys_t	vl_vdev_phys;				/* 112K	*/
	char		vl_uberblock[VDEV_UBERBLOCK_RING];	/* 128K	*/
} vdev_label_t;

typedef struct nvlist {
	int32_t		nvl_version;
	uint32_t	nvl_nvflag;	/* persistent flags */
	uint64_t	nvl_priv;	/* ptr to private data if not packed */
	uint32_t	nvl_flag;
	int32_t		nvl_pad;	/* currently not used, for alignment */
} nvlist_t;

typedef struct cksum_record {
	zio_cksum_t cksum;
	boolean_t labels[VDEV_LABELS];
	avl_node_t link;
} cksum_record_t;

typedef struct label {
        vdev_label_t label;
        nvlist_t *config_nv;
        cksum_record_t *config;
        cksum_record_t *uberblocks[MAX_UBERBLOCK_COUNT];
        boolean_t header_printed;
        boolean_t read_failed;
} label_t;
