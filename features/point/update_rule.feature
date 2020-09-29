Feature: 更新积分规则

	Background:
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
	Scenario: 商户更新自定义积分规则

		# jobs初始验证
		Given jobs登录系统
		When jobs添加积分规则
		"""
		{
			"type": "trade",
			"name": "rule1",
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
			"name": "rule2",
			"point": 30,
			"data": {
				"count": 314,
				"products": []
			}
		}
		"""

		Then jobs能获得积分规则列表
		"""
		[{
			"type": "money",
			"name": "rule2",
			"point": 30,
			"data": {
				"count": 314,
				"products": []
			},
			"is_system_rule": false
		}, {
			"type": "trade",
			"name": "rule1",
			"point": 20,
			"data": {
				"count": 2
			},
			"is_system_rule": false
		}, {
			"type": "valid_days",
			"name": "积分有效期",
			"data": {},
			"point": 0,
			"is_system_rule": true
		}, {
			"type": "max_per_day",
			"name": "积分获取上限",
			"data": {},
			"point": 0,
			"is_system_rule": true
		}, {
			"type": "protect_days",
			"name": "积分保护期",
			"data": {},
			"point": 0,
			"is_system_rule": true
		}]
		"""

		# jobs更新规则rule2
		When jobs更新积分规则'rule2'
		"""
		{
			"point": 9,
			"data": {
				"count": 5000,
				"products": [1,2]
			}
		}
		"""
		Then jobs能获得积分规则列表
		"""
		[{
			"type": "money",
			"name": "rule2",
			"point": 9,
			"data": {
				"count": 5000,
				"products": [1,2]
			},
			"is_system_rule": false
		}, {
			"type": "trade",
			"name": "rule1",
			"point": 20,
			"data": {
				"count": 2
			},
			"is_system_rule": false
		}, {
			"type": "valid_days",
			"name": "积分有效期",
			"data": {},
			"point": 0,
			"is_system_rule": true
		}, {
			"type": "max_per_day",
			"name": "积分获取上限",
			"data": {},
			"point": 0,
			"is_system_rule": true
		}, {
			"type": "protect_days",
			"name": "积分保护期",
			"data": {},
			"point": 0,
			"is_system_rule": true
		}]
		"""

		# jobs更新规则rule1
		When jobs更新积分规则'rule1'
		"""
		{
			"point": 11,
			"data": {
				"count": 111
			}
		}
		"""
		Then jobs能获得积分规则列表
		"""
		[{
			"type": "money",
			"name": "rule2",
			"point": 9,
			"data": {
				"count": 5000,
				"products": [1,2]
			},
			"is_system_rule": false
		}, {
			"type": "trade",
			"name": "rule1",
			"point": 11,
			"data": {
				"count": 111
			},
			"is_system_rule": false
		}, {
			"type": "valid_days",
			"name": "积分有效期",
			"data": {},
			"point": 0,
			"is_system_rule": true
		}, {
			"type": "max_per_day",
			"name": "积分获取上限",
			"data": {},
			"point": 0,
			"is_system_rule": true
		}, {
			"type": "protect_days",
			"name": "积分保护期",
			"data": {},
			"point": 0,
			"is_system_rule": true
		}]
		"""

	@ginger-crm @point
	Scenario: 商户更新系统积分规则

		# jobs初始验证
		Given jobs登录系统
		Then jobs能获得积分规则列表
		"""
		[{
			"type": "valid_days",
			"name": "积分有效期",
			"data": {},
			"point": 0,
			"is_system_rule": true
		}, {
			"type": "max_per_day",
			"name": "积分获取上限",
			"data": {},
			"point": 0,
			"is_system_rule": true
		}, {
			"type": "protect_days",
			"name": "积分保护期",
			"data": {},
			"point": 0,
			"is_system_rule": true
		}]
		"""

		# jobs更新系统积分
		When jobs更新积分规则'积分有效期'
		"""
		{
			"point": 1,
			"data": {
				"type": "days",
				"count": 10
			}
		}
		"""
		When jobs更新积分规则'积分获取上限'
		"""
		{
			"point": 2,
			"data": {
				"count": 200
			}
		}
		"""
		When jobs更新积分规则'积分保护期'
		"""
		{
			"point": 3,
			"data": {
				"count": 30
			}
		}
		"""
		Then jobs能获得积分规则列表
		"""
		[{
			"type": "valid_days",
			"name": "积分有效期",
			"data": {
				"type": "days",
				"count": 10
			},
			"point": 1,
			"is_system_rule": true
		}, {
			"type": "max_per_day",
			"name": "积分获取上限",
			"data": {
				"count": 200
			},
			"point": 2,
			"is_system_rule": true
		}, {
			"type": "protect_days",
			"name": "积分保护期",
			"data": {
				"count": 30
			},
			"point": 3,
			"is_system_rule": true
		}]
		"""

		# bill不受影响
		Given bill登录系统
		Then bill能获得积分规则列表
		"""
		[{
			"type": "valid_days",
			"name": "积分有效期",
			"data": {},
			"point": 0,
			"is_system_rule": true
		}, {
			"type": "max_per_day",
			"name": "积分获取上限",
			"data": {},
			"point": 0,
			"is_system_rule": true
		}, {
			"type": "protect_days",
			"name": "积分保护期",
			"data": {},
			"point": 0,
			"is_system_rule": true
		}]
		"""



