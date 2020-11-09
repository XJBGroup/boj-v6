<!-- group_doc -->
# Api Group Documentation ({{api_tag}})

!!insert_field:desc_{{api_tag}}!!

## Apis

{{body}}
## Local Object Reference

{{object_reference}}
<!-- group_body_doc -->
### {{operationId}}

The uri/restful key of this method is `{{key}}`

!!insert_field:desc_{{operationId}}!!

{{params_body}}
<!-- group_params_body_doc -->
+ `{{name}}`: {{type}}{{required}}: {{desc}}
    !!insert_field:desc_{{operationId}}_{{name}}!!

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
