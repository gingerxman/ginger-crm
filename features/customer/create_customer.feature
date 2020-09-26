Feature: 创建客户

@ginger-crm @customer @wip
Scenario: 用户登录创建客户
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

	# jobs初始验证
	Given jobs登录系统
	Then jobs能获得客户列表
	"""
	[]
	"""

	# lucy验证
	Given lucy注册为App用户
	Given lucy访问'jobs'的商城
	Given lucy访问'jobs'的商城
	Given lucy访问'jobs'的商城
	Then lucy能获得自己的客户信息
	"""
	{
		"name": "lucy"
	}
	"""

	# lily验证
	Given lily注册为App用户
	Given lily访问'jobs'的商城
	Then lily能获得自己的客户信息
	"""
	{
		"name": "lily"
	}
	"""

	# jobs验证
	Given jobs登录系统
	Then jobs能获得客户列表
	"""
	[{
		"name": "lily"
	}, {
		"name": "lucy"
	}]
	"""
	Then jobs能获得'lucy'的客户信息
	"""
	{
		"name": "lucy"
	}
	"""
	Then jobs能获得'lily'的客户信息
	"""
	{
		"name": "lily"
	}
	"""

	# bill验证
	Given bill登录系统
	Then bill能获得客户列表
	"""
	[]
	"""
	# lucy成为bill corp的客户
	Given lucy访问'bill'的商城
	Then lucy能获得自己的客户信息
	"""
	{
		"name": "lucy"
	}
	"""
	# bill验证
	Given bill登录系统
	Then bill能获得客户列表
	"""
	[{
		"name": "lucy"
	}]
	"""
