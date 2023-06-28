Feature: 更新角色
    測試更新角色相關的usecase功能

    Scenario: 成功更新角色
        Given 提供 <id>, <name>, <description>, <permissions>
        When 嘗試更新角色
        Then 角色應該成功被更新

        Examples:
            |id|name|description|permissions|
            |role1|new_name|new_description|["read:all", "write:all"]|

    Scenario: 提供的角色ID不存在
        Given 提供不存在的角色ID
        When 嘗試更新角色
        Then 應該返回一個錯誤，說明角色ID不存在


    Scenario: 提供的權限格式不正確
        Given 提供 <id>, <name>, <description>, <permissions>
        When 嘗試更新角色
        Then 應該返回一個錯誤，說明權限格式不正確

        Examples:
            |id|name|description|permissions|
            |role1|new_name|new_description|["read:all", "write:all", "invalid:all"]|

    Scenario: 提供的角色名稱或描述長度為最大值
        Given 提供 <id>, <name>, <description>, <permissions>
        When 嘗試更新角色
        Then 角色應該成功被更新

        Examples:
            |id|name|description|permissions|
            |role1|eMAWSxvuWc36VAKVFxMmeYHmr70GvI|rfgkzDeNnc69zIDnsxdZTfEazl2sXEfCKhFds6ydEWfzN5pGrRlQa22524xvzLS7gtgKFzqizI4aCxXIB7Vni2uPbWjy4vBntNc9XvnSKvAfqzbMOgmD3jxmKuJGNRO4zfX6HNykFQJfSB4qCu47bE6Uzhzul1uHXcrKQWRR85ziXcHMfu1g4NmMHQBpWiFswexTwn4g|["read:all"]|

    Scenario: 提供的角色名稱或描述長度超過最大值
        Given 提供 <id>, <name>, <description>, <permissions>
        When 嘗試更新角色
        Then 應該返回一個錯誤，說明角色名稱或描述長度超過最大值

        Examples:
            |id|name|description|permissions|
            |role1|o8SA8aJMn1EaBjMS3l6UPdLPZ931T9|QTdYDISBUy7YFfrPHAA9R34GHEPmotoGkT9k0JPXbZk5P2vk1WudZbwhVk2KrtgDrRPK9uaPcryLIFdBVL6l4ct2SdyBq7WI0htPXinMhjACuaN7x6RL7rhAeS3Esa6h9kNPoB3mAsFkzr9ysCFnlLajh8a0KlkJcAplKYvPOXVbrnEJ3mdfH3rzIaCxjCM4FU69K7hte|read:all|
            |role1|o8SA8aJMn1EaBjMS3l6UPdLPZ931T9c|QTdYDISBUy7YFfrPHAA9R34GHEPmotoGkT9k0JPXbZk5P2vk1WudZbwhVk2KrtgDrRPK9uaPcryLIFdBVL6l4ct2SdyBq7WI0htPXinMhjACuaN7x6RL7rhAeS3Esa6h9kNPoB3mAsFkzr9ysCFnlLajh8a0KlkJcAplKYvPOXVbrnEJ3mdfH3rzIaCxjCM4FU69K7ht|read:all|