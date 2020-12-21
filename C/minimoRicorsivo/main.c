#include <stdio.h>

int minimoRic(int a[], int n);

int main() {
	int a[] = { 20, 40, 9, 10, 30};

	printf("Minimo: %d\n", minimoRic(a, 5));
	
	return 0;
}

int minimoRic(int a[], int n) {
	n--;
	
	if (n == 0) {
		return a[0];
	}

	int min = minimoRic(a, n);
	
	if (a[n] < min) {
		return a[n];
	}
	return min;
}
