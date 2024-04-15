---
layout: default
title: Default Project Visibility Should Not Be Public
parent: Enterprise Policies
grand_parent: GitLab Policies
---


## Default Project Visibility Should Not Be Public
policy name: default_project_visibility_is_public

severity: HIGH

### Description
Your server is configured by default that every new project is created public. This means that it will be accessible by any member and unauthenticated user. It is recommended to turn the default option to be private
 and let admins control public visibility explicitly when needed.



### Remediation
1. Press Settings -> General
2. Expand "Visibility and access controls" section
3. Under "Default project visibility" toggle "Private"
4. Press "Save Changes"


