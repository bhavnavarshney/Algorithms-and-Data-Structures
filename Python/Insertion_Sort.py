print('***** INSERTION SORT *********')

def insertionSort(arr):
    
    for j in range(1,len(arr)):
        key = arr[j]
        index = j-1
        while(index>=0 and arr[index]>key):
            arr[index+1] = arr[index]
            index-=1
        
        arr[index+1] = key
    return arr

n = int(input('Enter no of elements: '))
string_arr = input('Enter your array here : ').split(' ')
arr = [int(num) for num in string_arr]
arr = insertionSort(arr)
print("SORTED ARRAY")
print(arr)

def insertionSort(arr):
    
    for j in range(1,len(arr)):
        key = arr[j]
        index = j-1
        while(index>=0 and arr[index]>key):
            arr[index+1] = arr[index]
            index-=1
        
        arr[index+1] = key
    return arr

# Worst Case Complexity = O(n*n)
# Best Case Complexity = O(n)
# In-place Algorithm
# Used when there are few elements and most of them are sorted