#include "stdio.h"

void findPair(int arr[10], int sum) {
	int left, right,s;
	left = 0;
	right = 9;
	while (left < right) {
		s = arr[left]+arr[right];
		if (s == sum) {
			printf ("Pair (%d, %d) : ", left, right);
			break;
		}
		else 
		if (s < sum) {
			left++;
		}
		else
		if (s > sum) {
			right++;
		}
	}
	if (left > right){
		printf ("no pair found");
	}
}

int main() {
	int arr[10];
	int i=0, sum, left, right;
	scanf("%d", &sum);
	while (i < 10) {
		scanf("%d", &arr[i]);
		i++;
	}
	i=0;
	while (i < 10) {
		printf("%d,", arr[i]);
		i++;
	}
	findPair(arr, sum);
	return 0;

}

