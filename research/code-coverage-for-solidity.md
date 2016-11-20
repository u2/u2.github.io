## Solc测试覆盖率的问题
1. 首先测试覆盖率并不能完成保证合约的正确率执行，比如无法解决DAO中`reentrancy`的bug

## 实现
1. 通过`parse`合约，然后注入一些额外的`event`代码。只注入`modifer`和`function`即可。

## 实现和难度
1. 三元操作符
2. `.call`和`throw`也要记录`event`
3. 隐藏的else分支

链接：
https://blog.colony.io/code-coverage-for-solidity-eecfa88668c2#.qiex8x527
https://ariya.io/2012/09/the-hidden-trap-of-code-coverage
https://en.wikipedia.org/wiki/Linear_code_sequence_and_jump
