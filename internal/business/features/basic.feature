#TODO
Feature: Create, retrieve and delete a conf
    Create a new configuration, retrieve the data value of that configuration and delete it.
    As an API user
    I need to be able to request a new configuration

    Scenario: User follow all steps
        When New feature with name <feature> and content <info>
        Then A new conf must be stored in both cache and datasource
        Then Api send http request 200
        When New get request hits the api with name <feature>
        Then Only data value field must be returned with http code 200 <info>
        When New delete request hits the api with name <feature>
        Then Conf with name <feature> must be deleted both datasource and cache
        When New get request hits the api with name <feature>
        Then Api sends no content found with http response code 204  

    Examples:
        | feature   | info          | 
        | fe1       | ./fe1.yml     | 
        | fe2       | ./fe2.yml     | 
