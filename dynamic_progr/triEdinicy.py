def f(n):
    if n == 1:
        return 2
    elif n == 2:
        return 4
    elif n == 3:
        return 7
    
    dp = [0] * (n+1)
    dp[1] = 2
    dp[2] = 4
    dp[3] = 7
    for i in range(4, n+1):
        dp[i] = dp[i-1] + dp[i-2] + dp[i-3]
    
    return dp[n]

n = int(input())
print(f(n))
