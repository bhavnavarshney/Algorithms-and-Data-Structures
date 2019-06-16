print('***** HEAP SORT ******')

def heapify(arr,n,i):
    largest = i
    left = 2*i+1
    right = 2*i+2

    if(left<n and arr[left]>arr[largest]):
        largest = left
    if (right<n and arr[right]>arr[largest]):
        largest = right
    if(largest!=i):
        arr[i],arr[largest] = arr[largest],arr[i]
        heapify(arr,n,largest)

def heapSort(arr):
    n = len(arr)
    for i in range(n//2,-1,-1):
        heapify(arr,n,i)
    
    for i in range(n-1,0,-1):
        arr[0],arr[i] = arr[i],arr[0]
        heapify(arr,i,0)

if __name__=='__main__':
    n = int(input('Enter the no. of elements: '))
    string_arr = input('Enter your array here : ').split(' ')
    arr = [int(num) for num in string_arr]
    heapSort(arr)
    print('***** SORTED ARRAY ******')
    print(arr)


# Total Time Complexity = O(nlogn)
# In-place Algorithm
# Heap Sort is less widely used as Merge and Quicksort are better in practice
# However, Heap Data Structure is enormously used