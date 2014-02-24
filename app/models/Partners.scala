/**
 * Fetch partners from the google doc
 *
 * @author Josh Fyne
 */

package models

import play.api.libs.ws.{WS,Response}
import play.api.libs.concurrent.Promise
import play.api.Play
import play.api.Play.{configuration,current}

case class Partner(name:String, link:String, logo:String)

object Partner {

    /**
     * Get a copy of the document
     *
     */
    def getContent = WS.url(Play.configuration.getString("application.partners").get).get()


    /**
     * Parse sheet
     *
     */
    def parse(sheet:Response):Seq[Partner] = {
        var partners:scala.collection.mutable.Buffer[Partner] = scala.collection.mutable.Buffer.empty[Partner]
        var lineNum = 0;
        sheet.body.split("\n").map(line => {
            if (lineNum != 0) {
                val data = line.split("\t")
                if (data.length == 3) {
                    partners += Partner(data(0), data(1), data(2))
                }
            }
            lineNum = lineNum + 1
        })
        partners.readOnly
    }
}
