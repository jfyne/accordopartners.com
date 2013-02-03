/**
 * Fetch expos from the google doc
 * 
 * @author Josh Fyne
 */

package models

import java.util.Date
import java.text.{SimpleDateFormat,ParsePosition}
import play.api.libs.ws.{WS,Response}
import play.api.libs.concurrent.Promise
import play.api.Play
import play.api.Play.{configuration,current}

case class Expo(name:String, date:Date, link:String, categoryName:String, categorySlug:String)

object Expos {

    /**
     * Get a copy of the document
     *
     */
    def getContent = WS.url(Play.configuration.getString("application.expos").get).get()
    
    /**
     * Parse sheet
     *
     */
    def parse(sheet:Response):Seq[Expo] = {
        val format = new SimpleDateFormat("dd/MM/yyyy")
        var expos:scala.collection.mutable.Buffer[Expo] = scala.collection.mutable.Buffer.empty[Expo]
        var lineNum = 0;
        sheet.body.split("\n").map(line => {
            if (lineNum != 0) {
                val data = line.split("\t")
                if (data.length == 5) {
                    expos += Expo(data(0), format.parse(data(1), new ParsePosition(0)), data(2), data(3), data(4))    
                }
            }
            lineNum = lineNum + 1
        })
        val now = new Date()
        now.setTime(now.getTime() - 86400000)
        expos.readOnly.filter(expo => expo.date.after(now))
    }

    /**
     * Convert Date to readable string
     *
     */
    def humanize(date:Date):String = {
        val dateOutput = new SimpleDateFormat("dd/MM/yyyy")
        dateOutput.format(date)
    }
}
