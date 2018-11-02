#include <inttypes.h>

struct pair {
	uint_fast64_t const first;
	uint_fast64_t const second;
};

static struct pair _fib(uint_fast64_t const n) {
	if(n == 0) {
		static struct pair const ret = {0, 1};
		return ret;
	}

	uint_fast64_t const a = _fib(n / 2).first;
	uint_fast64_t const b = _fib(n / 2).second;
	uint_fast64_t const c = a * (b * 2 - a);
	uint_fast64_t const d = a * a + b * b;

	if(n % 2 == 0) {
		return (struct pair const){c, d};
	} else {
		return (struct pair const){d, c + d};
	}
}

uint_fast64_t fib(uint_fast64_t const n) {
	return _fib(n).first;
}
