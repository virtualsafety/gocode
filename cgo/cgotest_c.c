// +build ignore
#include "cgotest_c.h"
#include <stdio.h>
void sofunc(char* str ){
	printf("%s\n",str);
	extern void goCallback(void);
	goCallback();
}
