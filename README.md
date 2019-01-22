Microservices for following endpoint:
/stock/{symbol}

API:http://localhost:8080/stock/MSFT,AAPL,HSBA.L?stock_exchange=NASDAQ,LSE

Method: GET

Description:

Reading a stocksymbol and  stock exchange from the user based on that getting response from  https://www.worldtradingdata.com
and Returns the stock prices of the given stock symbol from the exchanges provided.
symbol must be a valid stock symbol. 
If no stock_excahneg  is given then return the value of stock from AMEX (default) stock
exchange with not found error message.


