print('**** COUNTING SORT *******')

def countSort(arr):
    sorted = [0 for i in range(0,256)]
    count = [0 for i in range(0,256)]
    result = ["" for _ in arr]

    for i in arr:
        count[ord(i)] += 1
    #print(count)

    for i in range(256):
        count[i] += count[i-1]
    #print(count) 
    
    for i in range(len(arr)):
       sorted[count[ord(arr[i])]-1] = arr[i]
       count[ord(arr[i])] -= 1
    #print(sorted)

    for i in range(len(arr)):
        result[i] = sorted[i]
    return result



arr = "InformationTechnology"
ans = countSort(arr)
print('***** SORTED ARRAY ******')
print(ans)

#O(n+k); n = no of elements in input array and k is the input range