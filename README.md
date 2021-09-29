# NewPipe CSV to JSON converter

Golang solution to convert the youtube export .csv to the NewPipe import .json format  
Not sure why this needs to exist, but apparently it does.  
Also there are other people who made this in python etc, which I've found too late.   
https://github.com/JCGdev/Newpipe-CSV-Fixer   
there's also a website to do this whole thing which is probably far more convenient anyway:
https://juandarr.github.io/json-youtube-export/

# How to use

Follow NewPipe's guide on exporting your youtube subs, copied here for your convenience:

1. Go to this URL: https://takeout.google.com/takeout/custom/youtube
2. Log in when asked
3. Click on "All data included", then on "Deselect all", then select only "subscriptions" and click "OK"
4. Click on "Next step" and then on "Create export"
5. click on the "Download" button after it appears and
6. From the downloaded takeout zip extract the .csv file (usually under "YouTube and YouTube music/subscriptions/subscriptions.csv")

After you have that .csv file, run this script on it like so
```
go run main.go subscriptions.csv
```

you'll get a subscriptions.json that you should be able to import as a "PREVIOUS EXPORT" in NewPipe.
