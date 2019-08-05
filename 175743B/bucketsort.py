from multiprocessing import Pool
from time import time

MIN_BUCKET = 0
MAX_BUCKET = 10

def bucket_sort(aList):
    buckets = list()
    #print("---------------")
    for i in range(MIN_BUCKET, MAX_BUCKET + 1):
        buckets.append(None)
    for i in aList:
        buckets[i] = i
    #print(buckets)
    
    aList = [i for i in buckets if i is not None]
    
    #print(aList)
    return aList

def parallel_bucket_2_sort(aList):
    buckets = list()
    divid_n = MAX_BUCKET//2
    for i in range(MIN_BUCKET, divid_n + 1):
        buckets.append(None)
    for i in aList:
        if i < divid_n:
            buckets[i] = i
        else:
            buckets[i-divid_n-1] = i
    aList = [i for i in buckets if i is not None]
    
    #print(aList)
    return aList

if __name__ == '__main__':
    read_start = time()
    # ファイルをオープンする
    test_data = open("../ransu_10000000.txt", "r")
    
    # 行ごとにすべて読み込んでリストデータにする
    contents = test_data.readlines()
    
    #intに変換したものをnumbresにリストで渡す。
    numbers = [int(num) for num in contents]
    
    # ファイルをクローズする
    test_data.close()
    
    read_end= time()
    print("読み込み時間:",read_end - read_start)
    print("----------------")
    #データを分ける（４等分(0~numbers/4-1の範囲, numbers/4~...)）作業
    divid_buckets_start = time()
    MAX_BUCKET = max(numbers)
    #print("MAX_BUCKET",MAX_BUCKET)
    
    divided_number = MAX_BUCKET//2
    
    divided_list = [[],[]]
    
    for i in numbers:
        if i < divided_number:
            divided_list[0].append(i)
        else :
            divided_list[1].append(i)
    divid_buckets_end = time()
    print("リストを分けた時間:", divid_buckets_end - divid_buckets_start)
    print("----------------")
    
    #print("divided_list",divided_list) 
    
    simple_start=time()
    simple_list = bucket_sort(numbers)
    simple_end=time()
    #print("単純結果:",simple_list)
    print("単純時間:",simple_end - simple_start)
    print("----------------")
    
    parallel_start=time()
    p=Pool(2)
    parallel_list = p.map(parallel_bucket_2_sort, divided_list)
    p.close()
    parallel_list = [flatten for inner in parallel_list for flatten in inner]
    parallel_end=time()
    #print("並列結果:",parallel_list)
    print("並列時間:",parallel_end - parallel_start)