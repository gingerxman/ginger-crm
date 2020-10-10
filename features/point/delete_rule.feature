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
	Scenario: 商户删除积分规则

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

		# jobs删除规则rule2
		When jobs删除积分规则'rule2'
		Then jobs能获得积分规则列表
		"""
		[{
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

		# jobs不能删除系统规则
		When jobs删除积分规则'积分有效期'
		"""
		{
			"error": "point_rule:delete_fail"
		}
		"""
		Then jobs能获得积分规则列表
		"""
		[{
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

