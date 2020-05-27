#include "stdio.h"
#include "blkid/blkid.h"
#include "string.h"

int
main(int argc, char *argv[])
{
    char *devname = argv[1];

    char *type;

    char *tag = "TYPE";

    char *zfs_member = "zfs_member";

   type = blkid_get_tag_value(NULL, tag, devname);

   if (type != NULL) {
   	if (strcmp(type, zfs_member) == 0) {
   		printf ("In use by zfs (cstor or zfs localPV)");
   	}
   }
   else {
   	printf ("not in use by zfs");
   }

}

