typedef struct {
	int g, m, a;
} data;

int main() {
	data d;
	data *p;

	d.g = 21;
	d.m = 12;
	d.a = 2020;

	p = &d;
	p->g = 22;
	return 0;
}
