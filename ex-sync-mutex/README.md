# Example
We are going to take a simple use case here of storing, retreiving and deleting the data from a imaginary database. 

Imagine that there are multiple go routine accessing a critical section of te code, say saving a data, at the same time another routine tries to update some changes to the data or may be remove some data out of the existing data. If the save does not successfully done, then updating the data may fail or if update fails then the removal of some out of the existing data may get corrupted or throws exception. It may cause the data integrity issues. In this cases we generally use exclusive lock and in go we achieve that using sync package.

Here we shall explore this

DepositMoney(value int, chan status bool)

WithdrawMoney(value int, chan status bool)

