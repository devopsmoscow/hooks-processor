![DevOps Moscow Community](https://ucare.timepad.ru/ff573796-6d7c-4e68-8634-66295816dd6d/-/preview/308x600/-/format/jpeg/logo_org_126967.jpg)

# Hooks Processor

## NAME
hooks-processor - Process Dialogflow v2 webhooks with data captured from Telegram and forward it to DevOpsMoscow Bot services

## SYNOPSIS
**hooks-processor** COMMAND [OPTIONS]

## DESCRIPTION
**hooks-processor** is an app purposed to process data gathered from Telegram chat via Dialogflow webhook and send it to DevOpsMscow Bot backend services.

## COMMANDS
Available Commands:
  **help**   -     Help about any command
  **run**    -     Run hooks processing

## OPTIONS
The following options are understood:
  **--config string**  - config file (default is ./.config.yaml)
  **-h, --help**         -   help for run

## CONFIGURATION FILE
Configuration file is a *YAML* file by default located in the same directory as Hooks Processor binary. 
Example:

    port: 9000
    services:  
      list:  
        - service: meetup  
          actions:  
            - new_meetup  
            - get_meetup  
          url: http://meetup-svc/  
        - service: telegram  
          actions:  
            - get_members  
            - get_admins  
          url: http://telegram-svc

The following configuration properties are understood:

**port** - TCP port application will be listening for

**services** - list of structures which are defining URLs for downstream services

**services.list.service** - downstream service name

**services.list.url** - downstream service URL

**services.list.actions** - list of actions to be handled by particular downstream service
