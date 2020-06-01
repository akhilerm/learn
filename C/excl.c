#include <stdio.h>
#include <sys/types.h>
#include <sys/stat.h>
#include <fcntl.h>
#include <errno.h>
#include <unistd.h>
int main(int argc, char *argv[]) {

	int fd = open(argv[1], O_RDWR|O_EXCL);
	if (fd == -1) {
		if (errno == EBUSY)
			printf("device is in use\n");
		else
			perror("open failed:");
	} else {
		printf("Device is not in use");
		close(fd);
	}
	return 0;
}
