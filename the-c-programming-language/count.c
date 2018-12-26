#include <stdio.h>

int main(void)
{
    long nc;

    for(nc = 0; getchar() != EOF; ++nc)
        putchar("c");
    printf("%ld", nc);
}