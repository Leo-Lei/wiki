---
layout: post
title: Jackson samples
date: 2016-02-14 18:30:00
tags:
- Jackson
categories: Java
description: The tutoria will show you how to set up the Python environment.
---


# Convert json to object


# Convert map to object


# Custom JSON Deserialization with Jackson

# A full sample of converting List<Map> to List<>
```java
@JsonDeserialize(using = UserDeserializer.class)
public class User {
    private String name;
    private String email;
    private Gender gender;
    private Integer age;
    private Boolean isAdmin;

    public User(String name, String email, Gender gender, Integer age, Boolean isAdmin) {
        this.name = name;
        this.email = email;
        this.gender = gender;
        this.age = age;
        this.isAdmin = isAdmin;
    }

    public String getName() {
        return name;
    }

    public String getEmail() {
        return email;
    }

    public String getGender() {
        return gender;
    }

    public String getAge() {
        return age;
    }
    
    public Boolean getIsAdmin(){
      return isAdmin;
    }
}
```

```bash
NAME,EMAIL,GENDER,AGE,ISADMIN
```

```java
public class UserDeserializer extends JsonDeserializer<User> {
    @Override
    public User deserialize(JsonParser jsonParser, DeserializationContext deserializationContext) throws IOException, JsonProcessingException {

        ObjectCodec oc = jsonParser.getCodec();
        JsonNode node = oc.readTree(jsonParser);

        String name = node.get("NAME").asText();
        String email = node.get("EMAIL").asText();
        Gender gender = Gender.valueOf(node.get("GENDER").asText());
        Boolean isAdmin = node.get("ISADMIN").asBoolean();

        User user = new User(name, email, gender, isAdmin);
        
        return user;
    }
}

```

```java
NAME,EMAIL,GENDER,ISADMIN
```


```java

List<Map<String,Object>> users = new List<Map<String,Object>>();

Map<String,Object> user1 = new Map();
map.put("NAME","Jack");
map.put("EMAIL","jack@example.com");
map.put("GENDER","M");
map.put("ISADMIN","Y");

Map<String,Object> user2 = new Map();
map.put("NAME","Rose");
map.put("EMAIL","rose@example.com");
map.put("GENDER","F");
map.put("ISADMIN","N");

users.put(user1);
users.put(user2);

ObjectMapper mapper = new ObjectMapper();
List<User> userList = mapper.convertValue(users,mapper.getTypeFactory().constructCollectionType(List.class,User.class));

```



