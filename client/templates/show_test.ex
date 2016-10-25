ixrelay_parameterized (31c9250f-385d-408c-88a3-30a798bf48e7)
  description: a good ole test for some stuff
  bucket: DevInt (tlspa07hozys)
  created: 2016-04-11T12:13:48-04:00 by Devops (devops@icg360.com)
  trigger: https//api.runscope.com/radar/3bb01460-4842-4959-b2de-396c5fd43340/trigger
  lastrun: 2016-05-12T12:19:31-04:00 (0 errors)
  schedules: 5m, 15m
  steps:
    1. GET http://google.com
       note: test that ixconfig has some stuff
       assertions:
         response_status equal_number 200
         response_json.appname equal ixrelay
         response_json has_key version
       scripts:
         console.log("test)
    2. condition: backends not_empty
       note: test that ixconfig has some stuff
       steps:
         1. POST https://api.runscope.com/radar/3bb01460-4842-4959-b2de-396c5fd43340/batch?runscope_environment={{runscope_environment}}
            assertions:
              response_status equal_number 201
