#include <stdio.h>

void printArray(int a[], int n);
void sort(int a[], int n);

int main() {
	int a[] = {10, 12, 70, 22, 13, -2, 82, 21, 12};

	printArray(a, 9);

	sort(a, 9);
	printArray(a, 9);

	return 0;
}

void printArray(int a[], int n) {
	int i;

	printf("{  ");
	for (i = 0; i < n; i++) {
		printf("%d  ", a[i]);
	}
	printf("}\n");
}

void sort(int a[], int n) {
	int i, j, tmp, min, min_i;

	for (i = 0; i < n; i++) {
		min = a[i];
		min_i = i;
		for (j = i+1; j < n; j++) {
			if (a[j] < min) {
				min = a[j];
				min_i = j;
			}
		}
		tmp = a[i];
		a[i] = min;
		a[min_i] = tmp;
	}
}
