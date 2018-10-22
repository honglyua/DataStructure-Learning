# 快速排序
## 时间复杂度:: 
### 最坏情况下，即数组本身已经有序或存在多个相同值: O(n^2), 
### 平均情况下: O(nlogn), 
### 最好情况下: O(nlogn)
## 空间复杂度: O(1), 
## 是不稳定排序，何为稳定排序：
# 相比于归并排序，优先选择快速排序
# 算法：分治+递归
## 设置一个哨兵值，分别取出小于它的值进行排序，大于它的值进行排序。
## 重复上述过程，直到剩余一个元素
def quick_sort(a,lo,hi)
    if lo<hi
        # 设置哨兵值，返回数组中的位置
        p=partition(a,lo,hi)
        # 数组中lo到p的值为小于哨兵的值
        quick_sort(a,lo,p-1)
        # 数组中lo到p的值为大于哨兵的值
        quick_sort(a,p+1,hi)
    end
    return a
end
# 设置哨兵值，并返回数组中的位置
def partition(a,lo,hi)
    i=lo
    j=hi+1
    # 将第一个值设置为哨
    pivot= a[lo]
    while true
        #从左边开始循环找到第一个大于pivot的值，并终止循环
        begin
            i+=1
            break if i==hi
        end while a[i]<pivot

        ##从右边开始循环找到第一个小于pivot的值，并终止循环
        begin 
            j-=1
            break if j==lo
        end while a[j]>pivot

        # 跳出 while true循环，如果i>=j    
        break if i>=j

        #交换 arr[i]和arr[j]的值
        temp=a[i]
        a[i]=a[j]
        a[j]=temp
    end
    # 循环结束时，j的左边都是小于pivot的值，右边都是大于pivot的值
    # 最后交换arr[lo]和arr[j]的值
    temp=a[lo]
    a[lo]=a[j]
    a[j]=temp
    # 返回j
    return j
end

puts quick_sort([12,3,1,2,4,70,89,3,3],0,8)
# partition
# 第一次循环，交换 a[5] a[8]: 12,3,1,2,4,3,89,3,70,此时 i=5 j=8
# 第二次循环，交换 a[6] a[7]: 12,3,1,2,4,3,3,89,70,此时 i=6 j=7
# 第三次循环 i=7，j=6 终止循环，交换a[0],a[6]的值：3,3,1,2,4,3,12,89,70

# 空间复杂度：O(n)
def quick_sort_1(a)
    if a.length <= 1
        return a
    end
    # 随机选一个值为哨兵
    p = a.sample
    left = a.select{|e| e < p}
    right = a.select{|e| e > p}
    middle = a.select{|e| e==p}
    quick_sort_1(left) + middle + quick_sort_1(right)
end

puts quick_sort_1([12,3,1,2,4,70,89,3,3])