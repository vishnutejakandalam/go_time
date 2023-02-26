package main

import (
    "fmt"
    "sort"
    "sync"
)

func Sort(arr *[]int , wg *sync.WaitGroup) {
    defer wg.Done()
    sort.Ints(*arr)
}

func Partition(arr *[]int, parts int) []([]int) {
    avgSize := len(*arr) / parts
    result := make([]([]int), parts)
    nbLargerPartitions := len(*arr) - avgSize * parts

    from := 0
    for i := 0; i < parts; i++ {
        var offset int
        if i < nbLargerPartitions {
            offset = 1
        } else {
            offset = 0
        }

        to := from + (avgSize + offset)
        result[i] = (*arr)[from:to]

        from = to
    }

    return result
}

func Merge(partitions *[]([]int)) []int {
    resultSize := 0
    for _, partition := range *partitions {
        resultSize += len(partition)
    }

    result := make([]int, resultSize)
    currentIndexes := make([]int, len(*partitions))
    itemsProcessed := 0

    for itemsProcessed < resultSize {
        smallestIdx := 0
        smallestValue := int(^uint(0) >> 1)

        for idx, partition := range *partitions {
            if currentIndexes[idx] < len(partition) && partition[currentIndexes[idx]] < smallestValue {
                smallestValue = partition[currentIndexes[idx]]
                smallestIdx = idx
            }
        }

        result[itemsProcessed] = smallestValue
        currentIndexes[smallestIdx] += 1
        itemsProcessed += 1
    }

    return result
}

func main() {
    var wg sync.WaitGroup
    x := []int{28, 11, 93, 71, 59,  9, 95, 48, 95, 56, 3, 
                68, 83, 30, 23, 34, 87, 14, 36, 67, 64, 12, 
                17, 44, 28,  3, 25, 87, 74, 20, 80, 6,
                0, 87, 68, 56, 51, 84, 12, 82,  3, 90}

    partitions := Partition(&x, 4)

    wg.Add(4)
    for _, partition := range partitions {
        partitionCopy := partition
        go Sort(&partitionCopy, &wg)
    }

    wg.Wait()
    result := Merge(&partitions)
    fmt.Println(result)
}