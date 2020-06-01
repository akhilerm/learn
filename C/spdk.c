#include "stdint.h"
#include "linux/fs.h"
#include "stdio.h"
#include "sys/ioctl.h"
#include "fcntl.h"
#include "sys/types.h"
#include "unistd.h"
#include "stdlib.h"
#define _LARGEFILE64_SOURCE

#define SPDK_BLOBSTORE_TYPE_LENGTH 16
#define SPDK_BS_SUPER_BLOCK_SIG "SPDKBLOB"

typedef uint64_t spdk_blob_id;
typedef unsigned long long int var64;

struct spdk_bs_type {
	char bstype[SPDK_BLOBSTORE_TYPE_LENGTH];
};

struct spdk_bs_super_block {
	uint8_t		signature[8];
	uint32_t        version;
	uint32_t        length;
	uint32_t	clean; /* If there was a clean shutdown, this is 1. */
	spdk_blob_id	super_blob;

	uint32_t	cluster_size; /* In bytes */

	uint32_t	used_page_mask_start; /* Offset from beginning of disk, in pages */
	uint32_t	used_page_mask_len; /* Count, in pages */

	uint32_t	used_cluster_mask_start; /* Offset from beginning of disk, in pages */
	uint32_t	used_cluster_mask_len; /* Count, in pages */

	uint32_t	md_start; /* Offset from beginning of disk, in pages */
	uint32_t	md_len; /* Count, in pages */

	struct spdk_bs_type	bstype; /* blobstore type */

	uint32_t	used_blobid_mask_start; /* Offset from beginning of disk, in pages */
	uint32_t	used_blobid_mask_len; /* Count, in pages */

	uint64_t        size; /* size of blobstore in bytes */
	uint32_t        io_unit_size; /* Size of io unit in bytes */

	uint8_t         reserved[4000];
	uint32_t	crc;
};

int getLogicalBlockSize(int handle) {
        int lbSize = 0;

	if (ioctl(handle, BLKBSZGET, &lbSize)) {
		printf("getLogicalBlockSize: Reading LB size failed.\n");
		lbSize = 512;
        }
	printf("logical block size: %d\n", lbSize);

        return lbSize;
}

var64 readLBA(int handle, var64 lba, void* buf, var64 bytes) {
        int ret = 0;
        int lbSize = getLogicalBlockSize(handle);
        var64 offset = lba * lbSize;

        printf("readFromLBA: entered.\n");

        lseek64(handle, offset, SEEK_SET);
        ret = read(handle, buf, bytes);

	if(ret != bytes) {

		printf("read LBA: read failed.\n");
		return -1;
        }

        printf("read LBA: retval: %d.\n", ret);
        return ret;
}


int main(int argc, char *argv[]) {
	int fd;
  	struct spdk_bs_super_block *blob;
	blob=calloc(1, sizeof(*blob));

  	fd = open(argv[1], O_RDONLY);

  	if(fd == -1) {
    		printf("open %s failed", argv[1]);
    		exit(1);
  	}

  	memset(blob, 0, sizeof(*blob));
  	readLBA(fd, 0, blob, sizeof(*blob));
	
	printf ("LBA values/signature %s\n", blob->signature);
	
	if (memcmp(blob->signature, SPDK_BS_SUPER_BLOCK_SIG,
                   sizeof(blob->signature)) != 0) { 
                printf("Disk may not be in use by mayastor");
                return 0;
        } else {
		printf ("Disk in use by mayastor");
	}


	return 0;
}
