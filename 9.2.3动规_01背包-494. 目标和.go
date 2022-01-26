## 动规_01背包

给你一个整数数组 nums 和一个整数 target 。
向数组中的每个整数前添加 '+' 或 '-' ，然后串联起所有整数，可以构造一个 表达式 ：
例如，nums = [2, 1] ，可以在 2 之前添加 '+' ，在 1 之前添加 '-' ，然后串联起来得到表达式 "+2-1" 。
返回可以通过上述方法构造的、运算结果等于 target 的不同 表达式 的数目

输入：nums = [1,1,1,1,1], target = 3
输出：5
解释：一共有 5 种方法让最终目标和为 3 。
-1 + 1 + 1 + 1 + 1 = 3
+1 - 1 + 1 + 1 + 1 = 3
+1 + 1 - 1 + 1 + 1 = 3
+1 + 1 + 1 - 1 + 1 = 3
+1 + 1 + 1 + 1 - 1 = 3
即，让数组中的元素依次+或-，最后组成的表达式结果是目标值 元素可正可负

// 动规 二维
func findTargetSumWays(nums []int, target int) int {
    //先转化为数学问题
    //a-b=target
    //a+b=sum
    //a=(target+sum)/2
    //求出sum
    var sum int
    for _,value:=range nums{
        sum+=value
    }
    //如果sum<target或者 sum+target不是偶数（因为a是int） 或者两者之和小于0了
    if sum<target||(sum+target)%2==1||(sum+target)<0{
        return 0
    }
    //开始dp初始化
    dp:=make([][]int,len(nums)+1)
    for i:=0;i<=len(nums);i++{
        tmp:=make([]int,(target+sum)/2+1)//背包容量
        dp[i]=tmp
    }
    dp[0][0]=1//当背包容量为0，且物品为0时，填满背包就1种方法
    for i:=0;i<len(nums)+1;i++{
        if i==0{
            continue
        }
        for j:=0;j<(target+sum)/2+1;j++{
            if nums[i-1]<=j{//如果背包装的下
                dp[i][j]=dp[i-1][j]+dp[i-1][j-nums[i-1]]
            }else{
                dp[i][j]=dp[i-1][j]
            }
        }
    }
    return dp[len(nums)][(target+sum)/2]
}

// 一维 效率更高 内存消耗更少
func findTargetSumWays(nums []int, target int) int {
    sum:=0
	for i:=range nums{
		sum+=nums[i]
	}
	if target>sum {return 0}
	if (target+sum)%2==1{
		return 0
	}
	if target <0{
		if len(nums)==1 && sum< (-target) {return 0}
	}
	dp:=make([]int,(target+sum)/2+1)  //dp[j]+=dp[j-nums[i]]
	dp[0]=1
	for i:=0;i<len(nums);i++  {
		for j:=(target+sum)/2;j>=nums[i];j--{
			dp[j]+=dp[j-nums[i]]
		}
	}
	return dp[(target+sum)/2]
}
