Feature: 购买后获取积分

	Background:
		Given 系统配置虚拟资产
		"""
		[{
			"code": "point",
			"display_name": "积分",
			"exchange_rate": 1,
			"enable_fraction": false,
			"is_payable": true,
			"is_debtable": false
		}]
		"""

		Given ginger登录系统
		When ginger创建公司
		"""
		[{
			"name": "MIX",
			"username": "jobs"
		}, {
			"name": "BabyFace",
			"username": "bill"
		}]
		"""

	@ginger-crm @point
	Scenario: 1. 用户购买触发金额积分规则的商品
		# jobs创建商品
		Given jobs登录系统
		When jobs添加商品
		"""
		[{
			"name": "商品1",
			"price": 4.00
		}]
		"""
		When jobs添加积分规则
		"""
		{
			"type": "money",
			"name": "money_rule",
			"point": 30,
			"data": {
				"count": 314,
				"products": []
			}
		}
		"""

		# lucy的初始验证
		Given lucy访问'jobs'的商城
		Then lucy能获得虚拟资产'point'
		"""
		{
			"balance": 0
		}
		"""

		# lucy购买
		Given lucy访问'jobs'的商城
		When lucy购买'jobs'的商品
		"""
		{
			"products": [{
				"name": "商品1",
				"count": 2
			}]
		}
		"""
		When lucy支付最新订单
		Then lucy能获得虚拟资产'point'
		"""
		{
			"balance": 60
		}
		"""

	@ginger-crm @point
	Scenario: 2. 用户购买不能触发金额积分规则的商品
		# jobs创建商品
		Given jobs登录系统
		When jobs添加商品
		"""
		[{
			"name": "商品1",
			"price": 4.00
		}]
		"""
		When jobs添加积分规则
		"""
		{
			"type": "money",
			"name": "money_rule",
			"point": 30,
			"data": {
				"count": 514,
				"products": []
			}
		}
		"""

		# lucy的初始验证
		Given lucy访问'jobs'的商城
		Then lucy能获得虚拟资产'point'
		"""
		{
			"balance": 0
		}
		"""

		# lucy购买
		Given lucy访问'jobs'的商城
		When lucy购买'jobs'的商品
		"""
		{
			"products": [{
				"name": "商品1",
				"count": 1
			}]
		}
		"""
		When lucy支付最新订单
		Then lucy能获得虚拟资产'point'
		"""
		{
			"balance": 0
		}
		"""

	@ginger-crm @point
	Scenario: 3. 用户购买触发消费次数积分规则的商品:多次消费满足
		# jobs创建商品
		Given jobs登录系统
		When jobs添加商品
		"""
		[{
			"name": "商品1",
			"price": 4.00
		}]
		"""
		When jobs添加积分规则
		"""
		{
			"type": "trade",
			"name": "trade_rule",
			"point": 20,
			"data": {
				"count": 2
			}
		}
		"""

		# lucy的初始验证
		Given lucy访问'jobs'的商城
		Then lucy能获得虚拟资产'point'
		"""
		{
			"balance": 0
		}
		"""

		#lucy第一次购买
		Given lucy访问'jobs'的商城
		When lucy购买'jobs'的商品
		"""
		{
			"products": [{
				"name": "商品1",
				"count": 2
			}]
		}
		"""
		#lucy验证
		When lucy支付最新订单
		Then lucy能获得虚拟资产'point'
		"""
		{
			"balance": 0
		}
		"""

		#lucy第二次购买
		Given lucy访问'jobs'的商城
		When lucy购买'jobs'的商品
		"""
		{
			"products": [{
				"name": "商品1",
				"count": 2
			}]
		}
		"""
		#lucy验证
		When lucy支付最新订单
		Then lucy能获得虚拟资产'point'
		"""
		{
			"balance": 20
		}
		"""

		#lucy第三次购买
		Given lucy访问'jobs'的商城
		When lucy购买'jobs'的商品
		"""
		{
			"products": [{
				"name": "商品1",
				"count": 2
			}]
		}
		"""
		#lucy验证
		When lucy支付最新订单
		Then lucy能获得虚拟资产'point'
		"""
		{
			"balance": 20
		}
		"""

		#lucy第四次购买
		Given lucy访问'jobs'的商城
		When lucy购买'jobs'的商品
		"""
		{
			"products": [{
				"name": "商品1",
				"count": 2
			}]
		}
		"""
		#lucy验证
		When lucy支付最新订单
		Then lucy能获得虚拟资产'point'
		"""
		{
			"balance": 40
		}
		"""

	@ginger-crm @point
	Scenario: 4. 用户购买触发消费次数积分规则的商品:一次消费满足
		# jobs创建商品
		Given jobs登录系统
		When jobs添加商品
		"""
		[{
			"name": "商品1",
			"price": 4.00
		}]
		"""
		When jobs添加积分规则
		"""
		{
			"type": "trade",
			"name": "trade_rule",
			"point": 9,
			"data": {
				"count": 1
			}
		}
		"""

		# lucy的初始验证
		Given lucy访问'jobs'的商城
		Then lucy能获得虚拟资产'point'
		"""
		{
			"balance": 0
		}
		"""

		#lucy第一次购买
		Given lucy访问'jobs'的商城
		When lucy购买'jobs'的商品
		"""
		{
			"products": [{
				"name": "商品1",
				"count": 1
			}]
		}
		"""
		#lucy验证
		When lucy支付最新订单
		Then lucy能获得虚拟资产'point'
		"""
		{
			"balance": 9
		}
		"""

		#lucy第二次购买
		Given lucy访问'jobs'的商城
		When lucy购买'jobs'的商品
		"""
		{
			"products": [{
				"name": "商品1",
				"count": 1
			}]
		}
		"""
		#lucy验证
		When lucy支付最新订单
		Then lucy能获得虚拟资产'point'
		"""
		{
			"balance": 18
		}
		"""

	@ginger-crm @point
	Scenario: 5. 用户购买触发多条积分规则的商品
		# jobs创建商品
		Given jobs登录系统
		When jobs添加商品
		"""
		[{
			"name": "商品1",
			"price": 4.00
		}]
		"""
		When jobs添加积分规则
		"""
		{
			"type": "trade",
			"name": "trade_rule",
			"point": 20,
			"data": {
				"count": 2
			}
		}
		"""
		When jobs添加积分规则
		"""
		{
			"type": "money",
			"name": "money_rule",
			"point": 30,
			"data": {
				"count": 314,
				"products": []
			}
		}
		"""

		# lucy的初始验证
		Given lucy访问'jobs'的商城
		Then lucy能获得虚拟资产'point'
		"""
		{
			"balance": 0
		}
		"""

		#lucy第一次购买，只触发money rule
		Given lucy访问'jobs'的商城
		When lucy购买'jobs'的商品
		"""
		{
			"products": [{
				"name": "商品1",
				"count": 2
			}]
		}
		"""
		#lucy验证
		When lucy支付最新订单
		Then lucy能获得虚拟资产'point'
		"""
		{
			"balance": 60
		}
		"""

		#lucy第二次购买，同时触发money rule和trade rule
		Given lucy访问'jobs'的商城
		When lucy购买'jobs'的商品
		"""
		{
			"products": [{
				"name": "商品1",
				"count": 2
			}]
		}
		"""
		#lucy验证
		When lucy支付最新订单
		Then lucy能获得虚拟资产'point'
		"""
		{
			"balance": 140
		}
		"""

		#lucy第三次购买，只触发money_rule
		Given lucy访问'jobs'的商城
		When lucy购买'jobs'的商品
		"""
		{
			"products": [{
				"name": "商品1",
				"count": 2
			}]
		}
		"""
		#lucy验证
		When lucy支付最新订单
		Then lucy能获得虚拟资产'point'
		"""
		{
			"balance": 200
		}
		"""

		#lucy第四次购买，同时触发money rule和trade rule
		Given lucy访问'jobs'的商城
		When lucy购买'jobs'的商品
		"""
		{
			"products": [{
				"name": "商品1",
				"count": 2
			}]
		}
		"""
		#lucy验证
		When lucy支付最新订单
		Then lucy能获得虚拟资产'point'
		"""
		{
			"balance": 280
		}
		"""

	@ginger-crm @point @wip
	Scenario: 6. 用户购买触发系统积分规则:积分上限
		"""
		积分上限对money rule有效，对trade rule无效
		"""
		# jobs创建商品
		Given jobs登录系统
		When jobs添加商品
		"""
		[{
			"name": "商品1",
			"price": 4.00
		}]
		"""
		When jobs添加积分规则
		"""
		{
			"type": "money",
			"name": "money_rule",
			"point": 30,
			"data": {
				"count": 314,
				"products": []
			}
		}
		"""
		When jobs添加积分规则
		"""
		{
			"type": "trade",
			"name": "trade_rule",
			"point": 10,
			"data": {
				"count": 4,
				"products": []
			}
		}
		"""
		When jobs更新积分规则'积分获取上限'
		"""
		{
			"point": 0,
			"data": {
				"count": 75
			}
		}
		"""

		# lucy的初始验证
		Given lucy访问'jobs'的商城
		Then lucy能获得虚拟资产'point'
		"""
		{
			"balance": 0
		}
		"""

		#lucy第一次购买，不触发"积分上限"规则
		Given lucy访问'jobs'的商城
		When lucy购买'jobs'的商品
		"""
		{
			"products": [{
				"name": "商品1",
				"count": 2
			}]
		}
		"""
		#lucy验证
		When lucy支付最新订单
		Then lucy能获得虚拟资产'point'
		"""
		{
			"balance": 60
		}
		"""

		#lucy第二次购买，触发"积分上限"规则，积分累积至上限
		Given lucy访问'jobs'的商城
		When lucy购买'jobs'的商品
		"""
		{
			"products": [{
				"name": "商品1",
				"count": 2
			}]
		}
		"""
		#lucy验证
		When lucy支付最新订单
		Then lucy能获得虚拟资产'point'
		"""
		{
			"balance": 75
		}
		"""

		#lucy第三次购买，触发"积分上限"规则，积分无改变
		Given lucy访问'jobs'的商城
		When lucy购买'jobs'的商品
		"""
		{
			"products": [{
				"name": "商品1",
				"count": 2
			}]
		}
		"""
		#lucy验证
		When lucy支付最新订单
		Then lucy能获得虚拟资产'point'
		"""
		{
			"balance": 75
		}
		"""

		#lucy第四次购买，触发trade rule，"积分上限"规则不限制
		Given lucy访问'jobs'的商城
		When lucy购买'jobs'的商品
		"""
		{
			"products": [{
				"name": "商品1",
				"count": 2
			}]
		}
		"""
		#lucy验证
		When lucy支付最新订单
		Then lucy能获得虚拟资产'point'
		"""
		{
			"balance": 85
		}
		"""





