Feature: 取得角色列表
    測試取得角色列表相關的usecase功能

    Scenario: 成功獲取角色列表
        Given 提供 <limit> 和 <offset>
        When 嘗試獲取角色列表
        Then 應該成功獲取角色列表

        Examples:
            |limit|offset|
            |10|0|

    Scenario: 提供的limit或offset為負數
        Given 提供 <limit> 和 <offset>
        When 嘗試獲取角色列表
        Then 應該返回一個錯誤，說明limit或offset不能為負數

        Examples:
            |limit|offset|
            |-1|0|
            |10|-1|

    Scenario: 提供的limit為0
        Given 提供 <limit> 和 <offset>
        When 嘗試獲取角色列表
        Then 應該返回一個空的角色列表

        Examples:
            |limit|offset|
            |0|0|

    Scenario: 提供的offset超過實際角色數量
        Given 提供 <limit> 和 <offset>
        When 嘗試獲取角色列表
        Then 應該返回一個空的角色列表

        Examples:
            |limit|offset|
            |10|100|
