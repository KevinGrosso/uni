#include <stdio.h>

int sommaRic(int a[], int i, int n);
int sommaRic2(int a[], int n);

int main() {
	int a[] = { 10, 20, 30, 40};
	
	printf("Somma 1: %d\n",sommaRic(a, 0, 4));
	printf("Somma 2: %d\n",sommaRic2(a, 4));
	
	return 0;
}

int sommaRic(int a[], int i, int n) {
	if (i == n - 1) {
		return a[i];
	}
	return a[i] + sommaRic(a, i + 1, n);
}

int sommaRic2(int a[], int n) {
	n--;
	
	if (n == 0) {
		return a[0];
	}
	return a[n] + sommaRic2(a, n);
}
