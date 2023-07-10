Feature: 創建角色
    測試創建角色相關的usecase功能

    Scenario Outline: 創建角色成功
        
        Given 提供 <name>, <description>, <permissions>
        When 創建一個新角色
        Then 角色<name>, <description>, <permissions>，成功被創建
        
        Examples:
        	|name|description|permissions|
            |admin|this is admin role|read:all,write:all,delete:all|
            |develop|this is develop role|read:all|
            |owner|this is owner role|read:all,write:all|
            |test|this is test role|read:all,delete:all|

    Scenario Outline: 創建一個已存在的角色名稱
        
        Given 提供 <name>, <description>, <permissions>
        When 嘗試創建一個已存在的角色名稱
        Then 應該返回一個錯誤，說明角色名稱已存在
        
        Examples:
        	|name|description|permissions|
            |admin|this is admin role|read:all,write:all,delete:all|

    Scenario Outline: 提供的權限格式不正確
        
        Given 提供 <name>, <description>, <permissions>
        When 帶入有問題的輸入
        Then 應該返回一個錯誤，說明權限格式不正確
        
        Examples:
            |name|description|permissions|
            |invalid|this is invalid role|read:all,write:all,delete:all,invalid:all|

    Scenario Outline: 提供的名稱格式不正確
        
        Given 提供 <name>, <description>, <permissions>
        When 帶入有問題的輸入
        Then 應該返回一個錯誤，說明名稱格式不正確
        
        Examples:
            |name|description|permissions|
            |@|this is invalid role|read:all,write:all,delete:all|

    Scenario Outline: 提供的角色名稱或描述長度為最大值
            
        Given 提供 <name>, <description>, <permissions>
        When 創建一個新角色
        Then 角色<name>, <description>, <permissions>，成功被創建
        
        Examples:
            |name|description|permissions|
            |eMAWSxvuWc36VAKVFxMmeYHmr70GvI|rfgkzDeNnc69zIDnsxdZTfEazl2sXEfCKhFds6ydEWfzN5pGrRlQa22524xvzLS7gtgKFzqizI4aCxXIB7Vni2uPbWjy4vBntNc9XvnSKvAfqzbMOgmD3jxmKuJGNRO4zfX6HNykFQJfSB4qCu47bE6Uzhzul1uHXcrKQWRR85ziXcHMfu1g4NmMHQBpWiFswexTwn4g|read:all|

    Scenario Outline: 提供的角色描述長度超過最大值
        
        Given 提供 <name>, <description>, <permissions>
        When 帶入有問題的輸入
        Then 應該返回一個錯誤，說明角色描述長度超過最大值
        
        Examples:
            |name|description|permissions|
            |o8SA8aJMn1EaBjMS3l6UPdLPZ931T9|QTdYDISBUy7YFfrPHAA9R34GHEPmotoGkT9k0JPXbZk5P2vk1WudZbwhVk2KrtgDrRPK9uaPcryLIFdBVL6l4ct2SdyBq7WI0htPXinMhjACuaN7x6RL7rhAeS3Esa6h9kNPoB3mAsFkzr9ysCFnlLajh8a0KlkJcAplKYvPOXVbrnEJ3mdfH3rzIaCxjCM4FU69K7hte|read:all|
            |o8SA8aJMn1EaBjMS3l6UPdLPZ931T9|AQTdYDISBUy7YFfrPHAA9R34GHEPmotoGkT9k0JPXbZk5P2vk1WudZbwhVk2KrtgDrRPK9uaPcryLIFdBVL6l4ct2SdyBq7WI0htPXinMhjACuaN7x6RL7rhAeS3Esa6h9kNPoB3mAsFkzr9ysCFnlLajh8a0KlkJcAplKYvPOXVbrnEJ3mdfH3rzIaCxjCM4FU69K7ht|read:all|
            |o8SA8aJMn1EaBjMS3l6UPdLPZ931T9|CQTdYDISBUy7YFfrPHAA9R34GHEPmotoGkT9k0JPXbZk5P2vk1WudZbwhVk2KrtgDrRPK9uaPcryLIFdBVL6l4ct2SdyBq7WI0htPXinMhjACuaN7x6RL7rhAeS3Esa6h9kNPoB3mAsFkzr9ysCFnlLajh8a0KlkJcAplKYvPOXVbrnEJ3mdfH3rzIaCxjCM4FU69K7ht|read:all|

    Scenario Outline: 提供的角色名稱長度超過最大值
        
        Given 提供 <name>, <description>, <permissions>
        When 帶入有問題的輸入
        Then 應該返回一個錯誤，說明角色名稱長度超過最大值
        
        Examples:
            |name|description|permissions|
            |o8SA8aJMn1EaBjMS3l6UPdLPZ931T9c|QTdYDISBUy7YFfrPHAA9R34GHEPmotoGkT9k0JPXbZk5P2vk1WudZbwhVk2KrtgDrRPK9uaPcryLIFdBVL6l4ct2SdyBq7WI0htPXinMhjACuaN7x6RL7rhAeS3Esa6h9kNPoB3mAsFkzr9ysCFnlLajh8a0KlkJcAplKYvPOXVbrnEJ3mdfH3rzIaCxjCM4FU69K7ht|read:all|
            |o8SA8aJMn1EaBjMS3l6UPdLPZ93GT9c|QTdYDISBUy7YFfrPHAA9R34GHEPmotoGkT9k0JPXbZk5P2vk1WudZbwhVk2KrtgDrRPK9uaPcryLIFdBVL6l4ct2SdyBq7WI0htPXinMhjACuaN7x6RL7rhAeS3Esa6h9kNPoB3mAsFkzr9ysCFnlLajh8a0KlkJcAplKYvPOXVbrnEJ3mdfH3rzIaCxjCM4FU69K7ht|read:all|
            |o8SA8aJMn1EaBjMS3l6UPdLPZ93GT9casdfasdfasd|QTdYDISBUy7YFfrPHAA9R34GHEPmotoGkT9k0JPXbZk5P2vk1WudZbwhVk2KrtgDrRPK9uaPcryLIFdBVL6l4ct2SdyBq7WI0htPXinMhjACuaN7x6RL7rhAeS3Esa6h9kNPoB3mAsFkzr9ysCFnlLajh8a0KlkJcAplKYvPOXVbrnEJ3mdfH3rzIaCxjCM4FU69K7ht|read:all|