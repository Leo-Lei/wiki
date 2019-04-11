---
layout: post
title: Lua
date: 2016-06-17 14:50:00
tags:
- Windows
categories: Windows
---





```lua
boy = {money = 200}
function boy.goToMarket(self ,someMoney)
	self.money = self.money - someMoney
end
boy:goToMarket(100)
print(boy.money)
```
