---
layout: post
title: Mybatis
date: 2016-08-16 14:20:00
tags:
- Gradle
categories: 
- Java
- Gradle
---


```java
List<RuleDO> getRules(@Param("type") String type, @Param("name") String name);
```

```xml
<select id="getRules" resultMap="RuleMap">
        select
        id, name, type, service_name, service_type, enabled, metadata
        from rule
        <where>
            <if test="type != null and type != '' and type !='ALL'">
                type = #{type}
            </if>
            <if test="name != null and name != ''">
                <bind name="pattern" value="'%' + _parameter.name + '%'" />
                AND name like #{pattern}
            </if>
        </where>
    </select>
```
