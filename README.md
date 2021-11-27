## magazine

TODO 



### 操作


### Usage 

> go1.18
go install magazine

#### 初始化DB
./magazine db init --ensure

#### 

执行测试用例

ginkgo  -r 


### 缓存分级


###


#### 容灾措施


[key-set]
[key-value]


class Stock(
	getFormLocal()
	getFormCache()
	getFormStorge()
)

class StockRepo {
	lock
	getKeySet()
	unlock
}

class Magazine {
	lock()
	getKey()
	unlock
}

