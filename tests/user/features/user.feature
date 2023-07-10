Feature: 使用者 for usecase
    測試使用者相關的usecase功能

    @TestCaseKey=LA-T1
    Scenario: 註冊一個新用戶成功
        Given 註冊的使用者不存在
        When 註冊一個新用戶
        Then 用戶註冊成功

    @TestCaseKey=LA-T2
    Scenario: 註冊一個新用戶失敗
        Given 註冊的使用者已經存在
        When 註冊一個新用戶
        Then 用戶註冊失敗

    @TestCaseKey=LA-T3
    Scenario: 取得使用者個人資料
        Given 使用者已經登入
        When 取得使用者個人資料
        Then 取得使用者個人資料成功
