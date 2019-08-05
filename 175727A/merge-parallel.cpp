#include<iostream>
#include <random>
#include <time.h>
#include <omp.h>
#include <stdlib.h>
#define NUM_ITEMS 100000 //ソートする要素数
#define NUM_LIMIT 10000 //1要素の最大数
using namespace std;


void merge(int a[], int l, int m, int r)
{
	int temp[m-l+1], temp2[r-m];
	for(int i=0; i<(m-l+1); i++)
		temp[i]=a[l+i];
	for(int i=0; i<(r-m); i++)
		temp2[i]=a[m+1+i];
	int i=0, j=0, k=l;

	while(i<(m-l+1) && j<(r-m))
	{
		if(temp[i]<temp2[j])
			a[k++]=temp[i++];
		else
			a[k++]=temp2[j++];
	}

	while(i<(m-l+1))
		a[k++]=temp[i++];
	while(j<(r-m))
		a[k++]=temp2[j++];

}



void mergeSort(int a[], int l, int r)
{
	if(l<r)
	{
		int m=(l+r)/2;
		#pragma omp parallel sections num_threads(2)
		{
			#pragma omp section
			{
				mergeSort(a,l,m);
			}
			#pragma omp section
			{
				mergeSort(a,m+1,r);
			}
		}
		merge(a,l,m,r);
}


int main()
{
	int num,limit;
	int a[NUM_ITEMS];
	for (int i = 0; i < NUM_ITEMS; ++i)
		a[i]= random() % NUM_LIMIT+1;
	double st = omp_get_wtime();
	mergeSort(a,0,NUM_ITEMS-1);
	double en = omp_get_wtime();
	cout<<"計測時間:"<<en-st;
	cout<<"\n";
	return 0;
}
