Feature: Role assign usecase
    測試使用者指定角色的各種情況
    
    Scenario: 成功分配角色給用戶
        Given 提供 <id>, <userId>, <assign>
        When 嘗試分配角色給用戶
        Then 角色應該成功被分配給用戶

        Examples:
            |id|userId|assign|
            |role1|["user1", "user2"]|true|

    Scenario: 成功取消用戶的角色
        Given 提供 <id>, <userId>, <assign>
        When 嘗試取消用戶的角色
        Then 角色應該成功被從用戶那裡取消

        Examples:
            |id|userId|assign|
            |role1|["user1", "user2"]|false|

    Scenario: 提供的角色ID不存在
        Given 提供不存在的角色ID
        When 嘗試分配角色給用戶
        Then 應該返回一個錯誤，說明角色ID不存在

    Scenario: 提供的用戶ID不存在
        Given 提供不存在的使用者ID
        When 嘗試分配角色給用戶
        Then 應該返回一個錯誤，說明用戶ID不存在


    Scenario: 一次嘗試分配多個用戶到同一角色
        Given 提供 <id>, <userId>, <assign>
        When 嘗試分配角色給用戶
        Then 角色應該成功被分配給所有用戶

        Examples:
            |id|userId|assign|
            |role1|["user1", "user2", "user3", "user4", "user5"]|true|

    Scenario: 一次嘗試取消多個用戶的同一角色
        Given 提供 <id>, <userId>, <assign>
        When 嘗試取消用戶的角色
        Then 角色應該成功被從所有用戶那裡取消

        Examples:
            |id|userId|assign|
            |role1|["user1", "user2", "user3", "user4", "user5"]|false|