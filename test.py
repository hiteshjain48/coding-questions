arr = [1,-2,3,-4,5,7]
j = len(arr)-1
i=0
# for i in range(len(arr)):
#     if j<i:
#         break
#     if arr[i]<0 and arr[j]>0:
#         j-=1
#     if arr[i]>0 and arr[j]<0:
#         arr[i], arr[j] = arr[j],arr[i]
#         j-=1
while(j>i):
    if arr[i]<0 and arr[j]>0:
        j-=1
        i+=1
    elif arr[i]>0 and arr[j]<0:
        arr[i], arr[j] = arr[j],arr[i]
        j-=1
        i+=1
    elif arr[i] >0 and arr[j]>0:
        j-=1
    elif arr[i]<0 and arr[j]<0:
        i+=1
    print(arr)
    print(f"i={i}, j={j}\n")
print(arr)
