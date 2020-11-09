<!-- group_doc -->
# Api Group Documentation ({{api_tag}})

!!insert_field:desc_{{api_tag}}!!

## Apis

{{body}}
## Local Object Reference

{{object_reference}}
<!-- group_body_doc -->
### {{operationId}}

restful key: `{{key}}`

!!insert_field:desc_{{operationId}}!!

{{params_head}}
{{params_body}}
{{response_head}}
{{response_body}}

<!-- group_params_body_doc -->
+ `{{name}}`: {{type}}{{required}}: {{desc}}
    !!insert_field:desc_{{operationId}}_params_{{name}}!!

<!-- group_response_body_doc -->
+ code: `{{code}}`, type: {{type}}
    !!insert_field:desc_{{operationId}}_response_{{code}}_{{type}}!!

<!-- local_object_reference_doc -->

{{objects}}
<!-- object_doc -->
### [{{object_name}}]({{ref}})

+ type: {{type}}

+ fields:
    {{object_fields}}
    <!-- object_field_doc -->
    + `{{name}}`: {{type}}{{required}}: {{desc}}
        !!insert_field:desc_{{object_name}}_{{name}}!!
