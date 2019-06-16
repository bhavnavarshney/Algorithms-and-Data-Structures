print('**** OPTIMIZED BUBBLE SORT *******')

def bubbleSort(arr):
    n = len(arr)

    for i in range(0,n):
        swap = False
        for j in range(0,n-i-1):
            if(arr[j]>arr[j+1]):
                swap=True
                arr[j],arr[j+1] = arr[j+1],arr[j]
        if swap==False:
            break

if __name__=='__main__':
    n = int(input('Enter no. of elements: '))
    string_arr = input('Enter your array here: ').split(' ')
    arr = [int(num) for num in string_arr]
    bubbleSort(arr)
    print('***** SORTED ARRAY ******')
    print(arr)