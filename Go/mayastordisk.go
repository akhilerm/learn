// +build linux, cgo

package main


/*
 #cgo LDFLAGS: -lblkid
#include "string.h"
#include "stdlib.h"
#include "stdint.h"

#define SPDK_BLOBSTORE_TYPE_LENGTH 16

typedef uint64_t spdk_blob_id;

typedef struct {
        char bstype[SPDK_BLOBSTORE_TYPE_LENGTH];
}spdk_bs_type;

typedef struct  {
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
}spdk_bs_super_block;
*/
import "C"
import (
	"encoding/binary"
	"fmt"
	"io"
	"os"
)


func main()  {
	dev := os.Args[1]

	/*var spdkblob *C.char
	spdkblob = C.CString("SPDKBLOB")*/

	var spdk C.spdk_bs_super_block

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

	err = binary.Read(f, binary.BigEndian, &spdk)
	if err != nil {
		fmt.Println("error reading", err)
		return
	}

/*	var ptr *C.char
	ptr = (*C.char)(spdk.signature[0])
	s := C.GoString(ptr)*/

	fmt.Println(spdk.signature)

}

