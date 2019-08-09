#include <stdio.h>
#include <stdlib.h>
#include<iostream>
#include <random>
#include <time.h>
#include <omp.h>
#define NUM_ITEMS 10000000
#define NUM_LIMIT 100
using namespace std;

static void merge_sort(int* array, size_t size);
static void merge_sort_temp(int* array, int l, int r, int* temp);
static void merge(int* array, int l, int r, int m, int* temp);
static void print_array(const int* array, size_t size);

int main(void)
{
    int* array;
    array = (int *)malloc(sizeof(int) * NUM_ITEMS);
  	for (int i = 0; i < NUM_ITEMS; ++i)
  		array[i]= random() % NUM_LIMIT+1;

    double st = omp_get_wtime();
    merge_sort(array, NUM_ITEMS);
    double en = omp_get_wtime();
    cout<<"計測時間:"<<en-st;
  	cout<<"\n";
    free( array );
    return 0;
}

void merge_sort(int* array, size_t size)
{
    int* temp;
    temp = (int *)malloc(sizeof(int) * size);
    merge_sort_temp(array, 0, size-1, temp);
    free(temp);
}


void merge_sort_temp(int* array, int l, int r, int* temp)
{
    int m;
    if(l >= r){
        return;
    }

    m = (l+r) / 2;
    #pragma omp parallel sections num_threads(2)
		{
			#pragma omp section
			{
				merge_sort_temp(array, l, m, temp);
			}
			#pragma omp section
			{
				merge_sort_temp(array, m+1, r, temp);
			}
		}
    merge(array, l, r, m, temp);
}

void merge(int* array, int l, int r, int m, int* temp)
{
    int i, j, k;
    for(i = l; i <= m; ++i){
        temp[i] = array[i];
    }

    for(i = m+1, j = r; i <= r; ++i, --j){
        temp[i] = array[j];
    }

    i = l;
    j = r;
    for(k = l; k <= r; ++k) {
        if( temp[i] <= temp[j] ){
            array[k] = temp[i];
            ++i;
        }
        else{
            array[k] = temp[j];
            --j;
        }
    }
}
