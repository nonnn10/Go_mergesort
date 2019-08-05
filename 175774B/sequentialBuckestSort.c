#include <omp.h>
#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include <time.h>
#include "outputText.cpp"

#define dim 10000000 //要素数
//#define dim 4 //要素数

#define n_buckets 2 //バケツの数
struct bucket {
	int n_elem;
	int index; // [start : n_elem)
	int start; //starting point in B array
};

int cmpfunc (const void * a, const void * b)
{
 return ( *(int*)a - *(int*)b );
}

int main(int argc, char *argv[]) {

	int *A, *B, *temp;
	int  i, w, j, limit; //n_bucketsバケツの数
	struct bucket* buckets; //array of buckets
	double start_time,end_time,total_time;

	//printf("配列の要素数を入力してください \n");
	/* if (scanf("%d", &dim) != 1){
		printf("error\n");
		return -1;
	}*/
	/*/printf("バケツの数を入力してください \n");
	if (scanf("%d", &n_buckets) != 1){
		printf("error\n");
		return -1;
	}*/
	limit = 10000; //最大数
	w = (int)limit/n_buckets;

	//メモリ確保
	A = (int *) malloc(sizeof(int)*dim);
	B = (int *) malloc(sizeof(int)*dim);
	buckets = (struct bucket *) calloc(n_buckets, sizeof(struct bucket));

//乱数を生成格納
	for(i=0;i<dim;i++) {
		A[i] = random() % limit;
		//A[i] = random() % dim;

	}
	outputTextFile("unSorted.txt",A,dim);

// ****************************
// Starting the main algorithm
// ****************************

	start_time = omp_get_wtime();

//データ分割
	//#pragma omp parallel
	for (i=0; i<dim;i++){
		j = A[i]/w;
		if (j>n_buckets-1)
				j = n_buckets-1;
		buckets[j].n_elem++;
	}

	//buckets[0].index=0; //bucket 0 starts from 0 in B, bucket 1 starts from the start of bucket 0 + the number of elements in bucket 0 ...
	//buckets[0].start=0;
	//#pragma omp parallel for
	for (i=1; i<n_buckets;i++){
		buckets[i].index = buckets[i-1].index + buckets[i-1].n_elem;
		buckets[i].start = buckets[i-1].start + buckets[i-1].n_elem;
	}
	int b_index;
	//#pragma omp parallel for
	for (i=0; i<dim;i++){
		j = A[i]/w;
		if (j > n_buckets -1)
				j = n_buckets-1;
		b_index = buckets[j].index++;
		B[b_index] = A[i];
		//printf("%d,%d\n",B[b_index],b_index);
	}

	//#pragma omp parallel for
	for(i=0; i<n_buckets; i++)
		qsort(B+buckets[i].start, buckets[i].n_elem, sizeof(int), cmpfunc);

	temp = A;
	A = B;
	B = temp;
	end_time=omp_get_wtime();
	total_time = end_time - start_time;


	outputTextFile("Sorted.txt",A,dim);
	printf(" %d 個の要素をバブルソートするのに %f 秒\n", dim,total_time);

	/* int sorted = 1;

  	if (!sorted)
		printf("ソートされてません!!!\n");*/
}
