package normalize

var cmnPhrases = "H4sIAAAAAAAA/2yX+XKCyhLG/555ivuoGCOIiLtiAsY1ahbXxAio+DLM9ha3umeiOadOVaq6f983DgOzdfLYEuFavKzE4w/RkKe2ePyheVxRfoVAaGc0j33xspKLKcljP09tuZjSPCmr9ExYkqr0oulypyxPMn5c8uOSsGQPf47FHMs48tGWj7ZxNIDDSkWiA5I/Jzog9V2iA1LUQwozpMGM6GAojy3e3pM71I7GYbavrBH5C7/ObMGr72JxIv9i9NclogMQ7x0JhhVSgOPkAY6ThxHRASkusfWa/GYfoIkwzGMcvs5Q+8b+RTpB+tlp2iPFU3x27Qokd2OiA5B5F/MWqtAmOiA9HYgOSM820cEQS75Vr03uEHwbRzw7LB6Qv0Dz0ysr+TJwWX1KNDBryOpTdCpDObdVoUY0sJkvrw7Nsw7LjnJuy7NLNIBzuVD2MMmvO8LWP6x0pOzhmKdj/tKVgUvy80wDs4aUlT74NWS7HblllDlVaMecKrZwUmVHwjtrLc3jpvDO2rnIuU2Yc2EznzK/qPxM7IcEsk4q9kPK6j6r+zA+aIfAZj62RplowRC3BuJnZTSA0/fdUU797iinfnMWN3lBWQM05TRFp0pYA2SA9oay9jtrtIgOlIUv+PTwBZ8+TeW0blY5m/4op2rWOpstYMPC9E+Z/ap6bd5t8o0v1xXec8jNltcXsINvbbN6n/ccytYfrGvLRYtlB5FdyL+Ysu0jjmL7iKPYnnEtWeSWUdzstxNAE9vW87hpNA13R9nRv5z0mCcZ0YGy84F/dVE4j3RG2eXILleSn675JUTKlkQHyrIDO8+JDkDiUPxDsrBHmm6R3q9AclsEUsMvogPlxRKPXOEvlZ8RA96nag3/d4dOSrnd1EtsBYtLkxhYRhADi3LXgyMTQjuj3F3Ji2umDeHTTBt3M7OPuJuZTcQ9F3/rufhbJHZ8FUvPaBoo9yKYFO5FMCm8/WyOZe6vzbHMe46qzmGd9XcEoDOHddbfUd4753GPlxvklqGW1IgOSKc9W1fJLQONWR8onOdIpZ88tsgtQ811UCg/IXkO72OnOkOteWZti7Vd8hfQ6a1R87tIgznRAenbgZV8y1C7jsTrK7lloPHi0IxJZ6iV9yiU60jdHtEBabKU9Q3RGfNxhDwuyaXuAzPQROOL6ID0ZhMdkD4StvPILQNN1odEByD10CQ6ID0d2Aq/vc4oDzzYnBcXV4iBT1whgfdn7QDc1s4fJ088Vgn/4WuJ8n5NvKwI79fy1AbCNdOv4ZoJ4zxeSeedR+79UMyT+T9UfUDyaAD9zrq82CV/gfJoCQeWDpQPamBlR3LLQOPZTLkl0PjXRbUi0JRfEfM6aKqdiXlda5kqFlDrpKpYoHxZgWmIhuSWUb4v832L8L3L92NDvNzg0afRNIAj3nzQxGIDJMsJkCwnQOrUB1KnPuWHAssORAfKD++y+/b7zQ/v0h//fvMT7nx+asLOF+63cDPCKwfhZkByEQPJxZwKryEDVz7V2LJAhNdg1lADFXVf1EdqVIJdbmCcwF4XnWf4SBBaERWdL5gr0fmCuRLdCe87hG+6vO9Q8XTWl1ITvfDMg4mYpdj+F9B5WbH5A8lTm80fDPGqw2tVo2kAh0ch0YH+FoG6/NOEe6xotDytsWsRHOm93vu7Af2tEPU5pImVt0Zg5S39Ha0Z5+uLeHWIDlQsh+jNLfRWHaRVB2kzg4tGfBbhohHbMeutzRoXa4v5XbPSxa4quh9ErBJWblHxPYJ5E98jmDdpFXjJU5kPxzyAVVDZExzzd+ikVD709Tkuf76IgXYmf76orB6gO1k9YHeNT+gojx38VSuSvYToQGW3jVSJNb3Bq0h/DK8iuyu44GRjAbebDNz8NCXMGuanqaZXTa9ArPIGxCpvhnjvbIQ8raHmuyj4LjV3irlQgHonpN4JaeOw0gEFzO6autREGImvxd2UB1+EgdivoRV3PXC46yG1n5F87XVLSN2SIV5bymxlNIQXdHoJavA9ApdPBkiDMdKigLQoIL0NkaIUaTVDWs0M8bchj4pG41G