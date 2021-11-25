create table 

stock_01	10	20	29
stock_02 	12	12	12

get form stock_01 


get form stock_02


redis : 

	zset 
	
read stock_config 

register: stock 
[
	1-100: stock_01
	101-200:stock_02
	201-300:stock_03
]


read from config 
