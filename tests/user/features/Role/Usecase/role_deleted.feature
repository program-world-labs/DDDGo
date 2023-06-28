Feature: 刪除角色
    測試刪除角色相關的usecase功能

    Scenario: 成功刪除角色
        Given 提供 <id>
        When 嘗試刪除角色
        Then 角色應該成功被刪除

        Examples:
            |id|
            |role1|

    Scenario: 提供的角色ID不存在
        Given 提供不存在的角色ID
        When 嘗試刪除角色
        Then 應該返回一個錯誤，說明角色ID不存在


    Scenario: 嘗試刪除一個已經被分配給用戶的角色
        Given 提供一個已經被分配給用戶的角色
        When 嘗試刪除角色
        Then 應該返回一個錯誤，說明角色已經被分配給用戶
