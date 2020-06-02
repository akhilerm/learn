// +build linux, cgo

package main

/*
#include "string.h"
#include "stdlib.h"
#include "stdint.h"

#define SPDK_BLOBSTORE_TYPE_LENGTH 16

typedef uint64_t spdk_blob_id;

typedef struct {
        char bstype[SPDK_BLOBSTORE_TYPE_LENGTH];
}spdk_bs_type;

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
		spdk_bs_type     bstype;
		uint32_t        used_blobid_mask_start;
		uint32_t        used_blobid_mask_len;
		uint64_t        size;
		uint32_t        io_unit_size;
		uint8_t         reserved[4000];
		uint32_t        crc;
};

char *get_signature(struct spdk_bs_super_block *spdk)
{
	return spdk->signature;
}
*/

import "C"
import (
	"encoding/binary"
	"fmt"
	"io"
	"os"
)

const spdk_signature = "SPDKBLOB"


func main()  {
	dev := os.Args[1]
	var spdk *C.struct_spdk_bs_super_block
	buf := make([]byte, C.sizeof_struct_spdk_bs_super_block)
	f, err := os.Open(dev)
	if err != nil {
		fmt.Println("error opning", err)
		return
	}
	_, err = f.Seek(0, io.SeekStart)
	if err != nil {
		fmt.Println("error seeking", err)
		return
	}
	err = binary.Read(f, binary.LittleEndian, buf)
	if err != nil {
		fmt.Println("error reading", err)
		return
	}
	spdk = (*C.struct_spdk_bs_super_block)(C.CBytes(buf))
	var ptr *C.char
	ptr = (*C.char)(C.get_signature(spdk))
	s := C.GoString(ptr)

	// take only the first 8 characters for signature
	s = s[0:8]

	if s == spdk_signature {
		fmt.Println("in use by mayastor")
	} else {
		fmt.Println("not in use by mayastor")
	}

}