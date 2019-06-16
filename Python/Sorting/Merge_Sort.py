print('******** MERGE SORT *********')

def mergeSort(arr):
    if len(arr)>1:
        mid = len(arr)//2
        Left = arr[:mid]
        Right = arr[mid:]

        mergeSort(Left)
        mergeSort(Right)

        i=j=k=0
        while i<len(Left) and j<len(Right):
            if Left[i]<Right[i]:
                arr[k] = Left[i]
                i+=1
            else:
                arr[k] = Right[j]
                j+=1
            k+=1
        
        while i<len(Left):
            arr[k] = Left[i]
            i+=1
            k+=1
        
        while j<len(Right):
            arr[k] = Right[j]
            j+=1
            k+=1

        

if __name__=='__main__':
    n = int(input('Enter the no. of elements: '))
    string_arr = input('Enter your array here : ').split(' ')
    arr = [int(num) for num in string_arr]
    mergeSort(arr)
    print('**** SORTED ARRAY ****')
    print(arr)

# Worst, Average & Best case complexity is O(nlogn)
# Not an in-place Algorithm
# Used to sort linked list, External Sorting, Inversion Count Problem