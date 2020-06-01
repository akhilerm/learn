// +build linux, cgo

package main


/*
 #cgo LDFLAGS: -lblkid
#include "blkid/blkid.h"
#include "string.h"
#include "stdlib.h"

typedef uint64_t spdk_blob_id;
typedef unsigned long long int var64;

struct spdk_bs_type {
        char bstype[SPDK_BLOBSTORE_TYPE_LENGTH];
};

struct spdk_bs_super_block {
        uint8_t         signature[8];
        uint32_t        version;
        uint32_t        length;
        uint32_t        clean;
		spdk_blob_id    super_blob;

		uint32_t        cluster_size;

		uint32_t        used_page_mask_start;
		uint32_t        used_page_mask_len;

		uint32_t        used_cluster_mask_start;
		uint32_t        used_cluster_mask_len;

		uint32_t        md_start;
		uint32_t        md_len;

		struct spdk_bs_type     bstype;

		uint32_t        used_blobid_mask_start;
		uint32_t        used_blobid_mask_len;

		uint64_t        size;
		uint32_t        io_unit_size;

		uint8_t         reserved[4000];
		uint32_t        crc;
};
*/
import "C"
import (
	"fmt"
	"os"
	"unsafe"
)

const (
	ZFS_FILESYSTEM = "zfs_member"
	BLKID_FS_TYPE = "TYPE"
)

func main()  {

}

