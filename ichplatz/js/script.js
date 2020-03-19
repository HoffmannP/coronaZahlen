document.addEventListener('DOMContentLoaded', () => {
  window.fetch('coronaZahlen.json')
    .then(r => r.json())
    .then(render)
})

function render (data) {
  if (window.data === undefined) {
    window.data = data
  }
  window.data.Regions = Object.entries(data.Regions).map(([Name, region], i) => ({ Index: i, Name, ...region, Max: Math.max(region.Count, region.RKI, region.Mopo) }))
  window.data.Mopo.Count = window.data.Regions.map(r => r.Mopo).reduce((a, c) => a + c)
  document.querySelector('body').innerHTML = `
    <div class="container">
        <h1>Fallzahlen Corona Deutschlandweit</h1>

        <p>
            Die Fallzahlen sind aus den Homepages (z.T. Pressemitteilungen) der Bundesländer zusammengetragen.
            Die Gesamtzahl wird aus der Summe der für jedes Bundesland jeweils höchsten Fallzahl (Bundesland, 
            <a target="_blank" href="${data.RKI.URL}"><abbr title="Robert-Koch-Institut">RKI</abbr></a> bzw.
            <a target="_blank" href="${data.Mopo.URL}"><abbr title="Berliner Morgenpost">MoPo</abbr></a>) berechnet,
            in dem Bestreben, damit die aktuellsten Zahlen anzuzeigen.
        </p>
        <p>
            Diese Website ist ein privates Projekt. Der Quellcode steht unter der GPLv3 und ist auf 
            <a href="https://github.com/HoffmannP/coronaZahlen">Github</a> zu finden. Die Rohdaten sind separat im json-Format
            <a href="${data.Name}" download="coronaZahlen vom ${timestamp(data.Date)}.json}">downloadbar</a> herunterladbar. Kontakt zu
            mir gibt es via <a href="https://twitter.com/Hoffmann">Twitter&nbsp;<span class="tt">@Hoffmann</span></a>.
        </p>
        
        <table class="u-full-width">
            <thead>
                <tr>
                    <th class="sorted">Bundesland</th>
                    <th>Fälle</th>
                    <th>RKI</th>
                    <th>MoPo</th>
                    <th>Stand</th>
                </tr>
            </thead>
            <tbody>
                ${renderTable(data.Regions)}
            </tbody>
            <tfoot>
                <tr>
                    <th rowspan="3">Deutschland</th>
                    <th>${niceNumber(data.Sum)}</th>
                    <td colspan="2"></td>
                    <td>${timestamp(data.Date)}</td>
                </tr>
                <tr>
                    <td>RKI:</td>
                    <td>${niceNumber(data.RKI.Count)}</td>
                    <td></td>
                    <td>${timestamp(data.RKI.Date)}</td>
                </tr>
                <tr>
                    <td>Mopo:</td>
                    <td></td>
                    <td>${niceNumber(data.Mopo.Count)}</td>
                    <td>${timestamp(data.Mopo.Date)}</td>
                </tr>
            </tfoot>
        </table>
    </div>`
  window.th = document.querySelectorAll('thead > tr > th')
  window.th.forEach(
    (th, i) => th.addEventListener('click', sort.bind(0, i))
  )
  window.th[0].dataset.sort = 0
}

function renderTable (rows) {
  return rows.map(row => `
    <tr>
        <td><a href="${niceNumber(row.URL)}">${row.Name}</a></td>
        <td class="${row.Count < row.Max ? 'blass' : ''}">${isAvailable(row.Count)}</td>
        <td class="${row.RKI < row.Max ? 'blass' : ''}">${niceNumber(row.RKI)}</td>
        <td class="${row.Mopo < row.Max ? 'blass' : ''}">${niceNumber(row.Mopo)}</td>
        <td>${timestamp(row.Date)}</td>
    </tr>
  `).join('\n')
}

const sorter = [
  (a, b) => a.Index - b.Index,
  (b, a) => a.Index - b.Index,
  (a, b) => a.Count - b.Count,
  (b, a) => a.Count - b.Count,
  (a, b) => a.RKI - b.RKI,
  (b, a) => a.RKI - b.RKI,
  (a, b) => a.Mopo - b.Mopo,
  (b, a) => a.Mopo - b.Mopo,
  (a, b) => a.Date - b.Date,
  (b, a) => a.Date - b.Date
]

function sort (col) {
  window.th.forEach(th => th.classList.remove('sorted'))
  const th = window.th[col]
  th.classList.add('sorted')
  th.dataset.sort = +!+th.dataset.sort
  document.querySelector('tbody').innerHTML = renderTable(
    window.data.Regions.sort(sorter[2 * col + +th.dataset.sort])
  )
}

function timestamp (unix) {
  if (unix === 0) {
    return 'n/a'
  }
  const ts = new Date(unix * 1000)
  const withTime = ts.getHours() > 0 || ts.getMinutes() > 0
  const date = { day: 'numeric', month: '2-digit', year: 'numeric' }
  const time = { hour: 'numeric', minute: '2-digit' }
  return ts.toLocaleString('de-DE', { ...date, ...(withTime ? time : {}) }) + (withTime ? ' Uhr' : '')
}

const niceNumber = x => ('' + x).replace(/\B(?=(\d{3})+(?!\d))/g, '.')
const isAvailable = x => x === -1 ? 'n/a' : niceNumber(x)
