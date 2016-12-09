#include <stdio.h>

int main(int argc, char const *argv[])
{
	int i;
	int num;
	int sum = 0, avg = 0, min = 0, max = 0;

	for(i = 0; i < 3; i++)
	{
		printf("Please input the %dth number：\n", i + 1);
		scanf("%d", &num);
		if (0 == i)
		{
			min = num;
			max = num;
		}
		sum += num;
		if (num < min)
		{
			min = num;
		}
		if (num > max)
		{
			max = num;
		}
	}

	avg = sum / 3;

	printf("sum = %d, avg = %d, min = %d, max = %d\n", sum, avg, min, max);

	return 0;
}