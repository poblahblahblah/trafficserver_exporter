package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	//"io/ioutil"
	"net/http"
	"net/url"
	//"os"
	"reflect"
	"regexp"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/common/log"
	"github.com/prometheus/common/version"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	up = prometheus.NewDesc(
		"trafficserver_up",
		"Was talking to Trafficserver successfully",
		nil, nil,
	)
	invalidChars = regexp.MustCompile("[^a-zA-Z0-9:_]")
)

type TrafficServerCollector struct {
}

type Metrics struct {
	Counters Counters `json:"global"`
}

// Very incomplete list of counters, but these are the ones we know we care
// about right now. As we categorize and sort the metrics more, we'll bring
// more Counters and Gauges over from the structs folder.
type Counters struct {
	proxy_process_http_completed_requests                         float64 `json:"proxy.process.http.completed_requests"`
	proxy_process_http_total_incoming_connections                 float64 `json:"proxy.process.http.total_incoming_connections"`
	proxy_process_http_total_client_connections                   float64 `json:"proxy.process.http.total_client_connections"`
	proxy_process_http_total_server_connections                   float64 `json:"proxy.process.http.total_server_connections"`
	proxy_process_http_total_parent_proxy_connections             float64 `json:"proxy.process.http.total_parent_proxy_connections"`
	proxy_process_http_total_parent_retries                       float64 `json:"proxy.process.http.total_parent_retries"`
	proxy_process_http_total_parent_switches                      float64 `json:"proxy.process.http.total_parent_switches"`
	proxy_process_http_total_parent_retries_exhausted             float64 `json:"proxy.process.http.total_parent_retries_exhausted"`
	proxy_process_http_total_parent_marked_down_count             float64 `json:"proxy.process.http.total_parent_marked_down_count"`
	proxy_process_http_incoming_requests                          float64 `json:"proxy.process.http.incoming_requests"`
	proxy_process_http_outgoing_requests                          float64 `json:"proxy.process.http.outgoing_requests"`
	proxy_process_http_incoming_responses                         float64 `json:"proxy.process.http.incoming_responses"`
	proxy_process_http_invalid_client_requests                    float64 `json:"proxy.process.http.invalid_client_requests"`
	proxy_process_http_missing_host_hdr                           float64 `json:"proxy.process.http.missing_host_hdr"`
	proxy_process_http_get_requests                               float64 `json:"proxy.process.http.get_requests"`
	proxy_process_http_head_requests                              float64 `json:"proxy.process.http.head_requests"`
	proxy_process_http_trace_requests                             float64 `json:"proxy.process.http.trace_requests"`
	proxy_process_http_options_requests                           float64 `json:"proxy.process.http.options_requests"`
	proxy_process_http_post_requests                              float64 `json:"proxy.process.http.post_requests"`
	proxy_process_http_put_requests                               float64 `json:"proxy.process.http.put_requests"`
	proxy_process_http_push_requests                              float64 `json:"proxy.process.http.push_requests"`
	proxy_process_http_delete_requests                            float64 `json:"proxy.process.http.delete_requests"`
	proxy_process_http_purge_requests                             float64 `json:"proxy.process.http.purge_requests"`
	proxy_process_http_connect_requests                           float64 `json:"proxy.process.http.connect_requests"`
	proxy_process_http_extension_method_requests                  float64 `json:"proxy.process.http.extension_method_requests"`
	proxy_process_http_broken_server_connections                  float64 `json:"proxy.process.http.broken_server_connections"`
	proxy_process_http_cache_lookups                              float64 `json:"proxy.process.http.cache_lookups"`
	proxy_process_http_cache_writes                               float64 `json:"proxy.process.http.cache_writes"`
	proxy_process_http_cache_updates                              float64 `json:"proxy.process.http.cache_updates"`
	proxy_process_http_cache_deletes                              float64 `json:"proxy.process.http.cache_deletes"`
	proxy_process_http_tunnels                                    float64 `json:"proxy.process.http.tunnels"`
	proxy_process_http_throttled_proxy_only                       float64 `json:"proxy.process.http.throttled_proxy_only"`
	proxy_process_http_parent_proxy_transaction_time              float64 `json:"proxy.process.http.parent_proxy_transaction_time"`
	proxy_process_http_user_agent_request_header_total_size       float64 `json:"proxy.process.http.user_agent_request_header_total_size"`
	proxy_process_http_user_agent_response_header_total_size      float64 `json:"proxy.process.http.user_agent_response_header_total_size"`
	proxy_process_http_user_agent_request_document_total_size     float64 `json:"proxy.process.http.user_agent_request_document_total_size"`
	proxy_process_http_user_agent_response_document_total_size    float64 `json:"proxy.process.http.user_agent_response_document_total_size"`
	proxy_process_http_origin_server_request_header_total_size    float64 `json:"proxy.process.http.origin_server_request_header_total_size"`
	proxy_process_http_origin_server_response_header_total_size   float64 `json:"proxy.process.http.origin_server_response_header_total_size"`
	proxy_process_http_origin_server_request_document_total_size  float64 `json:"proxy.process.http.origin_server_request_document_total_size"`
	proxy_process_http_origin_server_response_document_total_size float64 `json:"proxy.process.http.origin_server_response_document_total_size"`
	proxy_process_http_parent_proxy_request_total_bytes           float64 `json:"proxy.process.http.parent_proxy_request_total_bytes"`
	proxy_process_http_parent_proxy_response_total_bytes          float64 `json:"proxy.process.http.parent_proxy_response_total_bytes"`
	proxy_process_http_pushed_response_header_total_size          float64 `json:"proxy.process.http.pushed_response_header_total_size"`
	proxy_process_http_pushed_document_total_size                 float64 `json:"proxy.process.http.pushed_document_total_size"`
	proxy_process_http_total_transactions_time                    float64 `json:"proxy.process.http.total_transactions_time"`
	proxy_process_http_cache_hit_fresh                            float64 `json:"proxy.process.http.cache_hit_fresh"`
	proxy_process_http_cache_hit_mem_fresh                        float64 `json:"proxy.process.http.cache_hit_mem_fresh"`
	proxy_process_http_cache_hit_revalidated                      float64 `json:"proxy.process.http.cache_hit_revalidated"`
	proxy_process_http_cache_hit_ims                              float64 `json:"proxy.process.http.cache_hit_ims"`
	proxy_process_http_cache_hit_stale_served                     float64 `json:"proxy.process.http.cache_hit_stale_served"`
	proxy_process_http_cache_miss_cold                            float64 `json:"proxy.process.http.cache_miss_cold"`
	proxy_process_http_cache_miss_changed                         float64 `json:"proxy.process.http.cache_miss_changed"`
	proxy_process_http_cache_miss_client_no_cache                 float64 `json:"proxy.process.http.cache_miss_client_no_cache"`
	proxy_process_http_cache_miss_client_not_cacheable            float64 `json:"proxy.process.http.cache_miss_client_not_cacheable"`
	proxy_process_http_cache_miss_ims                             float64 `json:"proxy.process.http.cache_miss_ims"`
	proxy_process_http_cache_read_error                           float64 `json:"proxy.process.http.cache_read_error"`
	proxy_process_http_err_client_abort_count_stat                float64 `json:"proxy.process.http.err_client_abort_count_stat"`
	proxy_process_http_err_client_abort_user_agent_bytes_stat     float64 `json:"proxy.process.http.err_client_abort_user_agent_bytes_stat"`
	proxy_process_http_err_client_abort_origin_server_bytes_stat  float64 `json:"proxy.process.http.err_client_abort_origin_server_bytes_stat"`
	proxy_process_http_err_connect_fail_count_stat                float64 `json:"proxy.process.http.err_connect_fail_count_stat"`
	proxy_process_http_err_connect_fail_user_agent_bytes_stat     float64 `json:"proxy.process.http.err_connect_fail_user_agent_bytes_stat"`
	proxy_process_http_err_connect_fail_origin_server_bytes_stat  float64 `json:"proxy.process.http.err_connect_fail_origin_server_bytes_stat"`
	proxy_process_http_misc_count_stat                            float64 `json:"proxy.process.http.misc_count_stat"`
	proxy_process_http_misc_user_agent_bytes_stat                 float64 `json:"proxy.process.http.misc_user_agent_bytes_stat"`
	proxy_process_http_http_misc_origin_server_bytes_stat         float64 `json:"proxy.process.http.http_misc_origin_server_bytes_stat"`
	proxy_process_http_background_fill_bytes_aborted_stat         float64 `json:"proxy.process.http.background_fill_bytes_aborted_stat"`
	proxy_process_http_background_fill_bytes_completed_stat       float64 `json:"proxy.process.http.background_fill_bytes_completed_stat"`
	proxy_process_http_cache_write_errors                         float64 `json:"proxy.process.http.cache_write_errors"`
	proxy_process_http_cache_read_errors                          float64 `json:"proxy.process.http.cache_read_errors"`
	proxy_process_http_100_responses                              float64 `json:"proxy.process.http.100_responses"`
	proxy_process_http_101_responses                              float64 `json:"proxy.process.http.101_responses"`
	proxy_process_http_1xx_responses                              float64 `json:"proxy.process.http.1xx_responses"`
	proxy_process_http_200_responses                              float64 `json:"proxy.process.http.200_responses"`
	proxy_process_http_201_responses                              float64 `json:"proxy.process.http.201_responses"`
	proxy_process_http_202_responses                              float64 `json:"proxy.process.http.202_responses"`
	proxy_process_http_203_responses                              float64 `json:"proxy.process.http.203_responses"`
	proxy_process_http_204_responses                              float64 `json:"proxy.process.http.204_responses"`
	proxy_process_http_205_responses                              float64 `json:"proxy.process.http.205_responses"`
	proxy_process_http_206_responses                              float64 `json:"proxy.process.http.206_responses"`
	proxy_process_http_2xx_responses                              float64 `json:"proxy.process.http.2xx_responses"`
	proxy_process_http_300_responses                              float64 `json:"proxy.process.http.300_responses"`
	proxy_process_http_301_responses                              float64 `json:"proxy.process.http.301_responses"`
	proxy_process_http_302_responses                              float64 `json:"proxy.process.http.302_responses"`
	proxy_process_http_303_responses                              float64 `json:"proxy.process.http.303_responses"`
	proxy_process_http_304_responses                              float64 `json:"proxy.process.http.304_responses"`
	proxy_process_http_305_responses                              float64 `json:"proxy.process.http.305_responses"`
	proxy_process_http_307_responses                              float64 `json:"proxy.process.http.307_responses"`
	proxy_process_http_3xx_responses                              float64 `json:"proxy.process.http.3xx_responses"`
	proxy_process_http_400_responses                              float64 `json:"proxy.process.http.400_responses"`
	proxy_process_http_401_responses                              float64 `json:"proxy.process.http.401_responses"`
	proxy_process_http_402_responses                              float64 `json:"proxy.process.http.402_responses"`
	proxy_process_http_403_responses                              float64 `json:"proxy.process.http.403_responses"`
	proxy_process_http_404_responses                              float64 `json:"proxy.process.http.404_responses"`
	proxy_process_http_405_responses                              float64 `json:"proxy.process.http.405_responses"`
	proxy_process_http_406_responses                              float64 `json:"proxy.process.http.406_responses"`
	proxy_process_http_407_responses                              float64 `json:"proxy.process.http.407_responses"`
	proxy_process_http_408_responses                              float64 `json:"proxy.process.http.408_responses"`
	proxy_process_http_409_responses                              float64 `json:"proxy.process.http.409_responses"`
	proxy_process_http_410_responses                              float64 `json:"proxy.process.http.410_responses"`
	proxy_process_http_411_responses                              float64 `json:"proxy.process.http.411_responses"`
	proxy_process_http_412_responses                              float64 `json:"proxy.process.http.412_responses"`
	proxy_process_http_413_responses                              float64 `json:"proxy.process.http.413_responses"`
	proxy_process_http_414_responses                              float64 `json:"proxy.process.http.414_responses"`
	proxy_process_http_415_responses                              float64 `json:"proxy.process.http.415_responses"`
	proxy_process_http_416_responses                              float64 `json:"proxy.process.http.416_responses"`
	proxy_process_http_4xx_responses                              float64 `json:"proxy.process.http.4xx_responses"`
	proxy_process_http_500_responses                              float64 `json:"proxy.process.http.500_responses"`
	proxy_process_http_501_responses                              float64 `json:"proxy.process.http.501_responses"`
	proxy_process_http_502_responses                              float64 `json:"proxy.process.http.502_responses"`
	proxy_process_http_503_responses                              float64 `json:"proxy.process.http.503_responses"`
	proxy_process_http_504_responses                              float64 `json:"proxy.process.http.504_responses"`
	proxy_process_http_505_responses                              float64 `json:"proxy.process.http.505_responses"`
	proxy_process_http_5xx_responses                              float64 `json:"proxy.process.http.5xx_responses"`
	proxy_process_https_incoming_requests                         float64 `json:"proxy.process.https.incoming_requests"`
	proxy_process_https_total_client_connections                  float64 `json:"proxy.process.https.total_client_connections"`
	proxy_process_http_origin_connections_throttled_out           float64 `json:"proxy.process.http.origin_connections_throttled_out"`
	proxy_process_http_post_body_too_large                        float64 `json:"proxy.process.http.post_body_too_large"`
	proxy_process_net_read_bytes                                  float64 `json:"proxy.process.net.read_bytes"`
	proxy_process_net_write_bytes                                 float64 `json:"proxy.process.net.write_bytes"`
	proxy_node_http_user_agents_total_documents_served            float64 `json:"proxy.node.http.user_agents_total_documents_served"`
	proxy_node_http_user_agents_total_transactions_count          float64 `json:"proxy.node.http.user_agents_total_transactions_count"`
	proxy_node_http_origin_server_total_transactions_count        float64 `json:"proxy.node.http.origin_server_total_transactions_count"`
	proxy_node_cache_bytes_total                                  float64 `json:"proxy.node.cache.bytes_total"`
}

func (c TrafficServerCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- up
}

func (c TrafficServerCollector) Collect(ch chan<- prometheus.Metric) {

	// FIXME: need to make these arguments
	var uri string = "http://localhost:8080/d6128c003f0179ad40d38cfc5a75b1e69b17145daaccfa02bf946983d2b6b9ea"
	var sslVerify bool = true
	var timeout = 5 * time.Second
	_, err := url.Parse(uri)

	if err != nil {
		ch <- prometheus.MustNewConstMetric(up, prometheus.GaugeValue, 0)
		return
	}

	body, err2 := fetchHTTP(uri, sslVerify, timeout)
	if err2 != nil {
		ch <- prometheus.MustNewConstMetric(up, prometheus.GaugeValue, 0)
		return
	}

	decoder := json.NewDecoder(body)
	var cont Metrics
	err3 := decoder.Decode(&cont)
	if err3 != nil {
		ch <- prometheus.MustNewConstMetric(up, prometheus.GaugeValue, 0)
		return
	}

	// This means things are healthy, so we can return an up
	ch <- prometheus.MustNewConstMetric(up, prometheus.GaugeValue, 1)

	// deal with all of our counters
	fields := reflect.TypeOf(cont.Counters)
	values := reflect.ValueOf(cont.Counters)
	num := fields.NumField()

	for i := 0; i < num; i++ {
		field := fields.Field(i)
		name := "trafficserver_" + invalidChars.ReplaceAllLiteralString(field.Name, "_")
		desc := prometheus.NewDesc(name, "Trafficserver metric "+field.Name, nil, nil)
		value := values.Field(i)
		ch <- prometheus.MustNewConstMetric(desc, prometheus.CounterValue, float64(value.Float()))
	}

	// do the same with the gauges, histograms, and summarys
	// TODO - figure out what metrics are gauges and which are counters
}

func fetchHTTP(uri string, sslVerify bool, timeout time.Duration) (io.Reader, error) {
	tr := &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: !sslVerify}}
	client := http.Client{
		Timeout:   timeout,
		Transport: tr,
	}

	resp, err := client.Get(uri)
	if err != nil {
		return nil, err
	}
	if !(resp.StatusCode >= 200 && resp.StatusCode < 300) {
		resp.Body.Close()
		return nil, fmt.Errorf("HTTP status %d", resp.StatusCode)
	}
	return resp.Body, nil
}

func main() {
	var (
		listenAddress = kingpin.Flag("web.listen-address", "Address to listen on for web interface and telemetry.").Default(":9101").String()
		metricsPath   = kingpin.Flag("web.telemetry-path", "Path under which to expose metrics.").Default("/metrics").String()
		//trafficServerScrapeURI    = kingpin.Flag("trafficserver.scrape-uri", "URI on which to scrape TrafficServer.").Default("http://localhost/stats").String()
		//trafficServerSSLVerify    = kingpin.Flag("trafficserver.ssl-verify", "Flag that enables SSL certificate verification for the scrape URI").Default("true").Bool()
		//trafficServerTimeout      = kingpin.Flag("trafficserver.timeout", "Timeout for trying to get stats from TrafficServer.").Default("5s").Duration()
	)

	log.AddFlags(kingpin.CommandLine)
	kingpin.Version(version.Print("trafficserver_exporter"))
	kingpin.HelpFlag.Short('h')
	kingpin.Parse()

	log.Infoln("Starting trafficserver_exporter", version.Info())
	log.Infoln("Build context", version.BuildContext())

	log.Infoln("Listening on", *listenAddress)
	http.Handle(*metricsPath, promhttp.Handler())
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`<html>
             <head><title>Trafficserver Exporter</title></head>
             <body>
             <h1>Trafficserver Exporter</h1>
             <p><a href='` + *metricsPath + `'>Metrics</a></p>
             </body>
             </html>`))
	})

	c := TrafficServerCollector{}
	prometheus.MustRegister(c)
	log.Fatal(http.ListenAndServe(*listenAddress, nil))
}
