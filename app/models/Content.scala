/**
 * Fetch content from the google doc
 *
 * @author Josh Fyne
 */

package models

import play.api.libs.ws.{WS,Response}
import play.api.libs.concurrent.Promise
import play.api.Play
import play.api.Play.{configuration,current}
import scala.io._

object Content {

    /**
     * Get a copy of the document
     *
     */
    def getContent = WS.url(Play.configuration.getString("application.content").get).get()


    /**
     * Parse sheet
     *
     */
    def parse(sheet:Response):Map[String,Seq[String]] = {
        var content:scala.collection.mutable.Map[String,Seq[String]] = scala.collection.mutable.Map.empty[String,Seq[String]]
        sheet.body.split("\n").map(line => {
            val pair = line.split("\t")
            if (pair.length > 0) content(pair(0)) = pair.tail
        })
        content.toMap
    }
}
